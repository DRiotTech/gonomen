package words_test

import (
	"testing"

	"github.com/DRiotTech/gonomen/words"
)

func TestGetEnglish(t *testing.T) {
	adjs, nouns, err := words.Get("en")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(adjs) < 50 {
		t.Errorf("expected at least 50 adjectives, got %d", len(adjs))
	}
	if len(nouns) < 50 {
		t.Errorf("expected at least 50 nouns, got %d", len(nouns))
	}
}

func TestGetUnsupportedLanguage(t *testing.T) {
	_, _, err := words.Get("xx")
	if err == nil {
		t.Error("expected error for unsupported language, got nil")
	}
}

func TestAllLanguages(t *testing.T) {
	langs := []string{"en", "el", "es", "it", "pl", "pt"}
	for _, lang := range langs {
		t.Run(lang, func(t *testing.T) {
			adjs, nouns, err := words.Get(lang)
			if err != nil {
				t.Fatalf("language %s: %v", lang, err)
			}
			if len(adjs) < 50 {
				t.Errorf("language %s: need 50+ adjectives, got %d", lang, len(adjs))
			}
			if len(nouns) < 50 {
				t.Errorf("language %s: need 50+ nouns, got %d", lang, len(nouns))
			}
		})
	}
}
