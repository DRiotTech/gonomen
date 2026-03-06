# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`gonomen` is a Go library (`github.com/nerdiken/gonomen`) for generating readable, multilingual usernames.

## Commands

```bash
# Build
go build ./...

# Test
go test ./...

# Run a single test
go test ./... -run TestName

# Test with coverage
go test ./... -cover
go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out

# Lint
golangci-lint run

# Vet
go vet ./...
```

## Architecture

**No `main` package** — this is a pure library. Entry point for users is `NewGenerator`.

### Core flow

```
NewGenerator(GeneratorOptions) → Generator
Generator.WithCase/WithLanguage/WithSuffixLength/WithSuffixType → Generator (immutable copy)
Generator.Generate() → string
```

### File layout

| File | Responsibility |
|------|---------------|
| `types.go` | `Case`, `SuffixType`, `GeneratorOptions` constants and types |
| `generator.go` | `Generator` struct, `NewGenerator`, `Generate` |
| `builder.go` | Fluent builder methods (`WithCase`, `WithLanguage`, `WithSuffixLength`, `WithSuffixType`) — each returns a new `Generator` copy |
| `formatter.go` | Internal `format(Case, adj, noun)` — applies casing (CamelCase, snake_case, kebab-case, etc.) |
| `suffix.go` | Internal `generateSuffix(SuffixType, length)` — produces digit or alphanumeric suffix |
| `words/words.go` | Registry + `Get(lang) (adjs, nouns, error)` |
| `words/*.go` | Word lists per language (`en`, `el`, `es`, `it`, `pl`, `pt`) — 100 adjectives + 100 nouns each |

### Key design decisions

- **Immutable builder**: every `WithXxx` method returns a new `Generator` value, never mutates the receiver.
- **Language fallback**: unsupported language codes silently fall back to `"en"`.
- **`SuffixLength: 0`** explicitly disables the suffix; `NewGenerator` defaults to 4 digits when the field is zero.
- **Word lists are ASCII-safe**: Greek (`el`) and Polish (`pl`) use transliterations so generated usernames are URL/login friendly.
- **Collision robustness**: 100 × 100 word pairs × 10,000 (4-digit suffix) = 100M combinations per language.
