package gonomen

// Case represents the casing format for generated usernames.
type Case string

const (
	CamelCase      Case = "camel"      // AdjectiveNoun
	LowerCamelCase Case = "lowerCamel" // adjectiveNoun
	SnakeCase      Case = "snake"      // adjective_noun
	KebabCase      Case = "kebab"      // adjective-noun
	Lower          Case = "lower"      // adjectivenoun
	Upper          Case = "upper"      // ADJECTIVENOUN
)

// SuffixType controls what characters are appended after the word pair.
type SuffixType string

const (
	SuffixDigits       SuffixType = "digits"       // e.g. 4821
	SuffixAlphanumeric SuffixType = "alphanumeric" // e.g. 4a2z
)

// Language identifies a supported word-list language (ISO 639-1 code).
// Supported values: "en", "el", "es", "it", "pl", "pt".
type Language string

// GeneratorOptions configures a Generator at creation time.
type GeneratorOptions struct {
	Language     Language   // default: EN
	Case         Case       // default: CamelCase
	SuffixLength int        // number of suffix characters; default 4
	SuffixType   SuffixType // default: SuffixDigits
}
