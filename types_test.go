package gonomen_test

import (
	"testing"

	"github.com/DRiotTech/gonomen"
)

func TestCaseConstants(t *testing.T) {
	cases := []struct {
		name string
		c    gonomen.Case
	}{
		{"CamelCase", gonomen.CamelCase},
		{"LowerCamelCase", gonomen.LowerCamelCase},
		{"SnakeCase", gonomen.SnakeCase},
		{"KebabCase", gonomen.KebabCase},
		{"Lower", gonomen.Lower},
		{"Upper", gonomen.Upper},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.c == "" {
				t.Errorf("case constant %s is empty", tc.name)
			}
		})
	}
}

func TestSuffixTypeConstants(t *testing.T) {
	types := []gonomen.SuffixType{gonomen.SuffixDigits, gonomen.SuffixAlphanumeric}
	for _, st := range types {
		if st == "" {
			t.Error("suffix type constant is empty")
		}
	}
}
