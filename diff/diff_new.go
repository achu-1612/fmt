package diff

import (
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Word-by-word comparison
func WordByWordDiff(text1, text2 string) string {
	dmp := diffmatchpatch.New()

	// Split the text into words for comparison
	diffs := dmp.DiffMain(strings.Join(strings.Fields(text1), " "), strings.Join(strings.Fields(text2), " "), false)
	return highlightDiff(diffs)
}

// Character-by-character comparison
func CharByCharDiff(text1, text2 string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)
	return highlightDiff(diffs)
}

// Line-by-line comparison
func LineByLineDiff(text1, text2 string) string {
	dmp := diffmatchpatch.New()
	// Split text into lines for comparison
	diffs := dmp.DiffMain(strings.Join(strings.Split(text1, "\n"), "\n"), strings.Join(strings.Split(text2, "\n"), "\n"), false)
	return highlightDiff(diffs)
}

// Helper function to apply HTML highlighting based on diff type
func highlightDiff(diffs []diffmatchpatch.Diff) string {
	var result strings.Builder

	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffDelete:
			result.WriteString(`<span class="deletion">` + diff.Text + `</span>`)
		case diffmatchpatch.DiffInsert:
			result.WriteString(`<span class="addition">` + diff.Text + `</span>`)
		case diffmatchpatch.DiffEqual:
			result.WriteString(diff.Text)
		}
	}
	return result.String()
}
