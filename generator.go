package gonomen

import (
	"math/rand/v2"

	"github.com/DRiotTech/gonomen/words"
)

const defaultSuffixLength = 4

// Generator generates random usernames according to its configuration.
// It is immutable: builder methods return a new copy.
type Generator struct {
	language     Language
	caseFormat   Case
	suffixLength int
	suffixType   SuffixType
}

// NewGenerator creates a Generator from the provided options.
// Zero values use defaults: language "en", CamelCase, 4 digit suffix.
func NewGenerator(opts GeneratorOptions) Generator {
	lang := opts.Language
	if lang == "" {
		lang = "en"
	}
	if _, _, err := words.Get(string(lang)); err != nil {
		lang = "en"
	}

	c := opts.Case
	if c == "" {
		c = CamelCase
	}

	sl := opts.SuffixLength
	if sl == 0 {
		sl = defaultSuffixLength
	}

	st := opts.SuffixType
	if st == "" {
		st = SuffixDigits
	}

	return Generator{
		language:     lang,
		caseFormat:   c,
		suffixLength: sl,
		suffixType:   st,
	}
}

// Generate returns a randomly generated username.
func (g Generator) Generate() string {
	adjs, nouns, _ := words.Get(string(g.language))
	adj := adjs[rand.IntN(len(adjs))]
	noun := nouns[rand.IntN(len(nouns))]
	return format(g.caseFormat, adj, noun) + generateSuffix(g.suffixType, g.suffixLength)
}
