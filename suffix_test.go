package gonomen

import (
	"strings"
	"testing"
	"unicode"
)

func TestGenerateSuffixDigits(t *testing.T) {
	s := generateSuffix(SuffixDigits, 4)
	if len(s) != 4 {
		t.Errorf("expected length 4, got %d", len(s))
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			t.Errorf("expected only digits, got %q", s)
		}
	}
}

func TestGenerateSuffixAlphanumeric(t *testing.T) {
	s := generateSuffix(SuffixAlphanumeric, 6)
	if len(s) != 6 {
		t.Errorf("expected length 6, got %d", len(s))
	}
	allowed := "abcdefghijklmnopqrstuvwxyz0123456789"
	for _, r := range s {
		if !strings.ContainsRune(allowed, r) {
			t.Errorf("unexpected character %q in suffix %q", r, s)
		}
	}
}

func TestGenerateSuffixZeroLength(t *testing.T) {
	s := generateSuffix(SuffixDigits, 0)
	if s != "" {
		t.Errorf("expected empty string for length 0, got %q", s)
	}
}

func TestGenerateSuffixUnknownTypeDefaultsToDigits(t *testing.T) {
	s := generateSuffix(SuffixType("unknown"), 4)
	for _, r := range s {
		if !unicode.IsDigit(r) {
			t.Errorf("unknown suffix type should default to digits, got %q", s)
		}
	}
}
