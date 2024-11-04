package diff

import "strings"

// CharDiff compares two strings and returns HTML-formatted string showing differences:
// - Characters present in text1 but not in text2 are marked in red (deletions)
// - Characters present in text2 but not in text1 are marked in green (additions)
// - Matching characters are left unchanged
func CharDiff(text1, text2 string) string {
	var result strings.Builder

	// Convert strings to rune slices to properly handle UTF-8 characters
	runes1 := []rune(text1)
	runes2 := []rune(text2)

	// Process the overlapping portion of both strings
	i := 0
	for i < len(runes1) && i < len(runes2) {
		if runes1[i] == runes2[i] {
			result.WriteRune(runes1[i])
		} else {
			result.WriteString(`<span class="deletion">` + string(runes1[i]) + `</span>`)
			result.WriteString(`<span class="addition">` + string(runes2[i]) + `</span>`)
		}
		i++
	}

	// Handle remaining characters in text1 (deletions)
	for ; i < len(runes1); i++ {
		result.WriteString(`<span class="deletion">` + string(runes1[i]) + `</span>`)
	}

	// Handle remaining characters in text2 (additions)
	for ; i < len(runes2); i++ {
		result.WriteString(`<span class="addition">` + string(runes2[i]) + `</span>`)
	}

	return result.String()
}

// wordDiff gets word-level differences.
func WordDiff(text1, text2 string) string {
	words1 := strings.Fields(text1)
	words2 := strings.Fields(text2)

	var result strings.Builder

	maxLen := len(words1)
	if len(words2) > maxLen {
		maxLen = len(words2)
	}

	for i := 0; i < maxLen; i++ {
		var word1, word2 string

		if i < len(words1) {
			word1 = words1[i]
		}

		if i < len(words2) {
			word2 = words2[i]
		}

		if word1 == word2 {
			result.WriteString(word1 + " ")
		} else {
			if word1 != "" {
				result.WriteString(`<span class="deletion">` + word1 + `</span>`)
			}

			if word2 != "" {
				result.WriteString(`<span class="addition">` + word2 + `</span>`)
			}
		}
	}

	return result.String()
}
