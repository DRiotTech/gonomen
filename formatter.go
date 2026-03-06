package gonomen

import "strings"

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func format(c Case, adj, noun string) string {
	switch c {
	case CamelCase:
		return capitalize(adj) + capitalize(noun)
	case LowerCamelCase:
		return strings.ToLower(adj) + capitalize(noun)
	case SnakeCase:
		return strings.ToLower(adj) + "_" + strings.ToLower(noun)
	case KebabCase:
		return strings.ToLower(adj) + "-" + strings.ToLower(noun)
	case Lower:
		return strings.ToLower(adj) + strings.ToLower(noun)
	case Upper:
		return strings.ToUpper(adj) + strings.ToUpper(noun)
	default:
		return capitalize(adj) + capitalize(noun)
	}
}
