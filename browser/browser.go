package browser

import (
	"syscall/js"
)

func add(a, b int) int {
	return a + b
}

func jsAdd(this js.Value, p []js.Value) interface{} {
	return js.ValueOf(add(p[0].Int(), p[1].Int()))
}

func registerCallbacks() {
	c := js.Global().Get("console")
	c.Call("log", "Go WebAssembly Initialized!")

	// Expose `add` function to JavaScript
	js.Global().Set("add", js.FuncOf(jsAdd))
}

func main() {
	c := make(chan struct{}, 0)

	// Register JS functions
	registerCallbacks()

	<-c // Block forever
}
