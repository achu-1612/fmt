//go:build js && wasm
// +build js,wasm

package diff

import (
	"strings"
	"syscall/js"
)

// charDiff gets character-level differences.
func charDiff(text1, text2 string) string {
	var result strings.Builder

	for i, c := range text1 {
		if i < len(text2) && c == rune(text2[i]) {
			result.WriteRune(c)
		} else {
			result.WriteString("<span style='color: red;'>" + string(c) + "</span>")
		}
	}

	for i := len(text1); i < len(text2); i++ {
		result.WriteString("<span style='color: green;'>" + string(text2[i]) + "</span>")
	}

	return result.String()
}

// wordDiff gets word-level differences.
func wordDiff(text1, text2 string) string {
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
				result.WriteString("<span style='color: red;'>" + word1 + "</span> ")
			}

			if word2 != "" {
				result.WriteString("<span style='color: green;'>" + word2 + "</span> ")
			}
		}
	}

	return result.String()
}

// Compare compares two texts and returns the differences.
func Compare(this js.Value, p []js.Value) interface{} {
	text1 := p[0].String()
	text2 := p[1].String()
	mode := p[2].String()

	var diffResult string

	if mode == "character" {
		diffResult = charDiff(text1, text2)
	} else {
		diffResult = wordDiff(text1, text2)
	}

	return js.ValueOf(diffResult)
}
