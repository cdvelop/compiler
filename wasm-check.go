package compiler

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Compiler) webAssemblyCheck() {
	// chequear si existe wasm_main.go en la ruta de trabajo ej: cmd/frontend/wasm_main.go

	fmt.Println("CARPETA DE TRABAJO: ", c.WORK_FOLDER, " ARCHIVO WASM: ", c.wasm_file_name)

	_, err := os.Open(filepath.Join(c.WORK_FOLDER, c.wasm_file_name))
	if err == nil {
		var wasm_compiler_name string

		c.wasm_build = true

		var remove_message string

		if c.with_tinyGo {

			// remove the message: syscall/js.finalizeRef not implemented
			//  https://github.com/tinygo-org/tinygo/issues/1140
			remove_message = `go.importObject.env["syscall/js.finalizeRef"] = () => {}`

			wasm_compiler_name = "TinyGo"

		} else {

			wasm_compiler_name = "Go"

		}

		c.js_wasm_import = `const go = new Go();
		` + remove_message + `
		WebAssembly.instantiateStreaming(fetch("static/app.wasm"), go.importObject).then((result) => {
			go.run(result.instance);
		});`

		fmt.Printf("*** Proyecto WebAssembly: [%v] Compilador: [%v] ***\n", c.wasm_file_name, wasm_compiler_name)
	}
}
