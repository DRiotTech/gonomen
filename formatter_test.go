package gonomen

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		c        Case
		adj      string
		noun     string
		expected string
	}{
		{"CamelCase", CamelCase, "happy", "river", "HappyRiver"},
		{"LowerCamelCase", LowerCamelCase, "happy", "river", "happyRiver"},
		{"SnakeCase", SnakeCase, "happy", "river", "happy_river"},
		{"KebabCase", KebabCase, "happy", "river", "happy-river"},
		{"Lower", Lower, "happy", "river", "happyriver"},
		{"Upper", Upper, "happy", "river", "HAPPYRIVER"},
		{"DefaultFallback", Case("unknown"), "happy", "river", "HappyRiver"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := format(tc.c, tc.adj, tc.noun)
			if got != tc.expected {
				t.Errorf("format(%s, %q, %q) = %q, want %q", tc.c, tc.adj, tc.noun, got, tc.expected)
			}
		})
	}
}
