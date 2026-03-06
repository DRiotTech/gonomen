package gonomen

import "github.com/nerdiken/gonomen/words"

// WithCase returns a new Generator with the given case format.
func (g Generator) WithCase(c Case) Generator {
	g.caseFormat = c
	return g
}

// WithLanguage returns a new Generator for the given language code.
// Falls back to "en" if the language is unsupported.
func (g Generator) WithLanguage(lang string) Generator {
	if _, _, err := words.Get(lang); err != nil {
		lang = defaultLanguage
	}
	g.language = lang
	return g
}

// WithSuffixLength returns a new Generator with the given suffix length.
// Pass 0 to disable the suffix entirely.
func (g Generator) WithSuffixLength(n int) Generator {
	g.suffixLength = n
	return g
}

// WithSuffixType returns a new Generator with the given suffix type.
func (g Generator) WithSuffixType(st SuffixType) Generator {
	g.suffixType = st
	return g
}
