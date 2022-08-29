package utils

import (
	"strings"
	"unicode"
)

func Systemify(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}

func Desystemify(input string) string {
	var builder strings.Builder
	words := strings.Split(input, "-")
	for _, w := range words {
		builder.WriteRune(rune(unicode.ToUpper(rune(w[0]))))
		builder.WriteString(w[1:] + " ")
	}
	return strings.TrimSuffix(builder.String(), " ")
}
