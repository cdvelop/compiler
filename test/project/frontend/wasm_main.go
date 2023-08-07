package main

import "syscall/js"

func main() {

	js.Global().Get("console").Call("log", "¡Hi 4 Go y WebAssembly!")
	select {}

}
