package gonomen_test

import (
	"strings"
	"testing"

	"github.com/nerdiken/gonomen"
)

func TestNewGeneratorDefaults(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{})
	result := g.Generate()
	if result == "" {
		t.Error("Generate() returned empty string")
	}
}

func TestGenerateHasSuffix(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{SuffixLength: 4})
	result := g.Generate()
	if len(result) < 4 {
		t.Fatalf("result too short: %q", result)
	}
	suffix := result[len(result)-4:]
	for _, r := range suffix {
		if r < '0' || r > '9' {
			t.Errorf("suffix %q contains non-digit %q", suffix, r)
		}
	}
}

func TestGenerateSnakeCase(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{
		Language:     "en",
		Case:         gonomen.SnakeCase,
		SuffixLength: 0,
	})
	result := g.Generate()
	if !strings.Contains(result, "_") {
		t.Errorf("snake_case result %q missing underscore", result)
	}
}

func TestGenerateItalian(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{Language: "it"})
	result := g.Generate()
	if result == "" {
		t.Error("Italian generator returned empty string")
	}
}

func TestGenerateUnsupportedLanguageFallsBackToEnglish(t *testing.T) {
	g := gonomen.NewGenerator(gonomen.GeneratorOptions{Language: "xx"})
	result := g.Generate()
	if result == "" {
		t.Error("Unsupported language should fallback to English, got empty")
	}
}

func TestGenerateAllLanguages(t *testing.T) {
	langs := []string{"en", "el", "es", "it", "pl", "pt"}
	for _, lang := range langs {
		t.Run(lang, func(t *testing.T) {
			g := gonomen.NewGenerator(gonomen.GeneratorOptions{Language: lang})
			result := g.Generate()
			if result == "" {
				t.Errorf("language %s: Generate() returned empty string", lang)
			}
		})
	}
}
