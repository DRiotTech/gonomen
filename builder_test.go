package gonomen_test

import (
	"strings"
	"testing"

	"github.com/nerdiken/gonomen"
)

func TestWithCase(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{SuffixLength: 0}).
		WithCase(gonomen.SnakeCase)
	result := g.Generate()
	if !strings.Contains(result, "_") {
		t.Errorf("WithCase(SnakeCase): expected underscore in %q", result)
	}
}

func TestWithCaseKebab(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{SuffixLength: 0}).
		WithCase(gonomen.KebabCase)
	result := g.Generate()
	if !strings.Contains(result, "-") {
		t.Errorf("WithCase(KebabCase): expected hyphen in %q", result)
	}
}

func TestWithLanguage(t *testing.T) {
	for _, lang := range []gonomen.Language{"es", "it", "pl", "pt", "el"} {
		t.Run(string(lang), func(t *testing.T) {
			g := gonomen.NewGenerator(gonomen.GeneratorOptions{}).WithLanguage(lang)
			result := g.Generate()
			if result == "" {
				t.Errorf("WithLanguage(%s) returned empty string", lang)
			}
		})
	}
}

func TestWithLanguageUnsupportedFallback(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{}).WithLanguage("zz")
	result := g.Generate()
	if result == "" {
		t.Error("unsupported language should fallback to English")
	}
}

func TestWithSuffixLength(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{}).WithSuffixLength(6)
	result := g.Generate()
	suffix := result[len(result)-6:]
	for _, r := range suffix {
		if r < '0' || r > '9' {
			t.Errorf("suffix %q contains non-digit %q", suffix, r)
		}
	}
}

func TestWithSuffixLengthZero(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{}).WithSuffixLength(0)
	result := g.Generate()
	if result == "" {
		t.Error("Generate() with suffix 0 should still return word pair")
	}
	// Should have no digits at end when suffix is 0
	for _, r := range result {
		if r >= '0' && r <= '9' {
			t.Errorf("expected no digits when SuffixLength=0, got %q", result)
			break
		}
	}
}

func TestWithSuffixType(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{SuffixLength: 8}).
		WithSuffixType(gonomen.SuffixAlphanumeric)
	result := g.Generate()
	if len(result) < 8 {
		t.Errorf("result too short: %q", result)
	}
}

func TestBuilderChaining(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{}).
		WithLanguage("it").
		WithCase(gonomen.KebabCase).
		WithSuffixLength(3).
		WithSuffixType(gonomen.SuffixAlphanumeric)
	result := g.Generate()
	if !strings.Contains(result, "-") {
		t.Errorf("expected kebab-case in %q", result)
	}
}

func TestBuilderImmutability(t *testing.T) {
	base := gonomen.NewGenerator(gonomen.GeneratorOptions{})
	snake := base.WithCase(gonomen.SnakeCase).WithSuffixLength(0)
	camel := base.WithCase(gonomen.CamelCase).WithSuffixLength(0)

	snakeResult := snake.Generate()
	camelResult := camel.Generate()

	if !strings.Contains(snakeResult, "_") {
		t.Errorf("snake generator gave %q, expected underscore", snakeResult)
	}
	if strings.Contains(camelResult, "_") {
		t.Errorf("camel generator gave %q, should not have underscore", camelResult)
	}
}
