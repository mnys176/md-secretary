package utils

import "strings"

func prepareInput(input string) []string {
	raw := strings.Split(strings.TrimSpace(input), "\n\n")
	clean := make([]string, 0)
	for _, r := range raw {
		s := strings.TrimSpace(r)
		if len(s) > 0 {
			clean = append(clean, s)
		}
	}
	return clean
}

func fitParagraphToLineLength(paragraph string, length int) string {
	var paragraphBuilder, lineBuilder strings.Builder
	words := strings.Split(paragraph, " ")

	// adjusts the cutoff length to account for linebreak behavior
	var specialOffset int

	for _, w := range words {
		if strings.Contains(w, "\n") {
			lineBuilder.WriteString(w + " ")
			paragraphBuilder.WriteString(lineBuilder.String())
			lineBuilder.Reset()

			// account for extra length added by the last word
			subwords := strings.Split(w, "\n")
			lastWord := subwords[len(subwords)-1]
			specialOffset = len(lastWord)
			continue
		}

		// if adding this word will exceed the desired length, start a new line
		if lineBuilder.Len()+len(w) > length-specialOffset {
			paragraphBuilder.WriteString(lineBuilder.String() + "\n")
			lineBuilder.Reset()

			// remove this side-effect if it was added
			specialOffset = 0
		}
		lineBuilder.WriteString(w + " ")
	}
	// add last line to the paragraph and return
	paragraphBuilder.WriteString(lineBuilder.String())
	return paragraphBuilder.String()
}

func ChopString(input string, length int) string {
	paragraphs := prepareInput(input)

	// truncate each paragraph to predefined line length
	var outputBuilder strings.Builder
	for _, p := range paragraphs {
		outputBuilder.WriteString(fitParagraphToLineLength(p, length) + "\n\n")
	}

	// trim trailing newlines
	return strings.TrimSuffix(outputBuilder.String(), "\n\n")
}
