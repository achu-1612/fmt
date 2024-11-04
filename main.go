//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/achu-1612/fmt/diff"
)

// compare compares two texts and returns the differences.
func compare(this js.Value, p []js.Value) interface{} {
	fmt.Println(this, p)
	if len(p) != 3 {
		return js.Error{Value: js.Global().Get("Error").New("expected 3 arguments")}
	}
	text1 := p[0].String()
	text2 := p[1].String()
	mode := p[2].String()

	var diffResult string

	if mode == "character" {
		diffResult = diff.CharByCharDiff(text1, text2)
	} else {
		diffResult = diff.WordByWordDiff(text1, text2)
	}

	return js.ValueOf(diffResult)
}

func registerCallbacks() {
	js.Global().Set("compare", js.FuncOf(compare))
}

func main() {
	c := make(chan struct{}, 0)

	registerCallbacks()

	<-c
}
