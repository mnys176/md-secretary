package utils

import "strings"

func snapToNearestMultiple(input int, multiple int) int {
	var lower, upper int
	for {
		if lower + multiple >= input {
			break
		}
		lower += multiple
	}
	for {
		upper += multiple
		if upper >= input {
			break
		}
	}
	if input-lower > upper-input {
		return upper
	}
	return lower
}

func SanitizeUnorderedList(raw string, expectedIndent int) string {
	var previousIndent int
	var builder strings.Builder
	prefix := "* "
	lines := strings.Split(strings.TrimSpace(raw), "\n")
	for _, l := range lines {
		// eliminate trailing whitespace while keeping left spaces
		l = strings.TrimRight(l, " ")
		l = strings.Trim(l, "\t\r")

		// separate raw indent and examine for potential adjustment
		spaces, text, hasPrefix := strings.Cut(l, prefix)
		currentIndent := len(spaces)
		if hasPrefix {
			// correct misaligned indent if it is slightly off
			if currentIndent % expectedIndent != 0 {
				currentIndent = snapToNearestMultiple(currentIndent, expectedIndent)
			}

			// children must always be one level higher than their parent
			if currentIndent > previousIndent + expectedIndent {
				currentIndent = previousIndent + expectedIndent
			}

			// scale indent down for consistency
			spaces = strings.Repeat(" ", currentIndent / expectedIndent)
			previousIndent = currentIndent
			builder.WriteString(spaces+prefix+text+"\n")
		}
	}

	// trim trailing newline
	return builder.String()[:builder.Len()-1]
}

func SanitizeParagraphs(raw string) string {
	var builder strings.Builder
	paragraphs := strings.Split(strings.TrimSpace(raw), "\n\n")
	for _, p := range paragraphs {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			builder.WriteString(p+"\n\n")
		}
	}

	// trim trailing newlines
	return builder.String()[:builder.Len()-2]
}

func SanitizeRegExp(raw string) string {
	var builder strings.Builder
	for _, r := range raw {
		if strings.ContainsRune(`\.+*?()|[]{}^$`, r) {
			builder.WriteRune('\\')
		}
		builder.WriteRune(r)
	}
	return builder.String()
}
