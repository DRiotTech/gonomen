# gonomen

[![CI](https://github.com/DRiotTech/gonomen/actions/workflows/ci.yml/badge.svg)](https://github.com/DRiotTech/gonomen/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/DRiotTech/gonomen/badge.svg?branch=main)](https://coveralls.io/github/DRiotTech/gonomen?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/DRiotTech/gonomen)](https://goreportcard.com/report/github.com/DRiotTech/gonomen)
[![Go Reference](https://pkg.go.dev/badge/github.com/DRiotTech/gonomen.svg)](https://pkg.go.dev/github.com/DRiotTech/gonomen)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.25-00ADD8?logo=go)](go.mod)
[![Repo Size](https://img.shields.io/github/repo-size/DRiotTech/gonomen)](https://github.com/DRiotTech/gonomen)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

A Go library for generating readable, multilingual usernames — like `HappyRiver4821` or `feliz-luna9k3`.

Combines a random adjective and noun from a curated word list, formats them in your preferred casing style, and appends an optional suffix. Zero external dependencies.

## Install

```sh
go get github.com/DRiotTech/gonomen
```

## Quick start

```go
import "github.com/DRiotTech/gonomen"

// Defaults: English, CamelCase, 4-digit suffix
g := gonomen.NewGenerator(gonomen.GeneratorOptions{})
fmt.Println(g.Generate()) // e.g. "HappyRiver4821"
```

## Options

```go
g := gonomen.NewGenerator(gonomen.GeneratorOptions{
    Language:     "it",                       // ISO 639-1 code; default "en"
    Case:         gonomen.KebabCase,          // default CamelCase
    SuffixLength: 6,                          // default 4; 0 to disable
    SuffixType:   gonomen.SuffixAlphanumeric, // default SuffixDigits
})
fmt.Println(g.Generate()) // e.g. "felice-mare9k3a2b"
```

### Fluent builder

The generator is immutable — each `With*` call returns a new copy.

```go
base := gonomen.NewGenerator(gonomen.GeneratorOptions{})

snake := base.WithCase(gonomen.SnakeCase).WithSuffixLength(0)
kebab := base.WithCase(gonomen.KebabCase).WithLanguage("es")

fmt.Println(snake.Generate()) // e.g. "happy_river"
fmt.Println(kebab.Generate()) // e.g. "feliz-luna4821"
```

## Casing formats

| Constant              | Example       |
|-----------------------|---------------|
| `CamelCase` (default) | `HappyRiver`  |
| `LowerCamelCase`      | `happyRiver`  |
| `SnakeCase`           | `happy_river` |
| `KebabCase`           | `happy-river` |
| `Lower`               | `happyriver`  |
| `Upper`               | `HAPPYRIVER`  |

## Suffix types

| Constant                 | Example |
|--------------------------|---------|
| `SuffixDigits` (default) | `4821`  |
| `SuffixAlphanumeric`     | `4a2z`  |

Pass `SuffixLength: 0` to generate a bare word pair with no suffix.

## Supported languages

| Code | Language                      |
|------|-------------------------------|
| `en` | English                       |
| `el` | Greek (ASCII transliteration) |
| `es` | Spanish                       |
| `it` | Italian                       |
| `pl` | Polish (ASCII transliteration)|
| `pt` | Portuguese                    |

Unsupported language codes silently fall back to English.

## Collision resistance

Each language ships 100 adjectives × 100 nouns. With the default 4-digit suffix that gives **100,000,000 unique combinations per language** — sufficient for most use cases without any uniqueness check on the caller side.

## License

MIT
