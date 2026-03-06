package words

import "fmt"

// WordList holds adjectives and nouns for a language.
type WordList struct {
	Adjectives []string
	Nouns      []string
}

var registry = map[string]WordList{}

// Get returns the adjective and noun lists for the given language code.
func Get(lang string) ([]string, []string, error) {
	wl, ok := registry[lang]
	if !ok {
		return nil, nil, fmt.Errorf("gonomen: unsupported language %q", lang)
	}
	return wl.Adjectives, wl.Nouns, nil
}
