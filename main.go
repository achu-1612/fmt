//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/achu-1612/fmt/diff"
)

func registerCallbacks() {
	js.Global().Set("compareTexts", js.FuncOf(diff.Compare))
}

func main() {
	c := make(chan struct{}, 0)

	registerCallbacks()

	<-c
}
