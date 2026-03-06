package gonomen

import (
	"math/rand/v2"
	"strings"
)

const (
	charsetDigits       = "0123456789"
	charsetAlphanumeric = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func generateSuffix(st SuffixType, length int) string {
	if length == 0 {
		return ""
	}
	charset := charsetDigits
	if st == SuffixAlphanumeric {
		charset = charsetAlphanumeric
	}
	var sb strings.Builder
	sb.Grow(length)
	for range length {
		sb.WriteByte(charset[rand.IntN(len(charset))])
	}
	return sb.String()
}
