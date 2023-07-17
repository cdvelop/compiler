package compiler

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Compiler) webAssemblyCheck() {
	// chequear si existe wasm_main.go en la ruta de trabajo ej: frontend/wasm_main.go
	_, err := os.Open(filepath.Join(c.WORK_FOLDER, c.wasm_file_name))
	if err == nil {
		var wasm_compiler_name string

		c.wasm_build = true

		if c.with_tinyGo {
			wasm_compiler_name = "TinyGo"
		} else {
			wasm_compiler_name = "Go"
		}

		fmt.Printf("*** Proyecto WebAssembly: [%v] Compilador: [%v] ***\n", c.wasm_file_name, wasm_compiler_name)
	}
}
