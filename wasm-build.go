package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (c Compiler) BuildWASM() {
	err := c.buildWASM(filepath.Join(c.WORK_FOLDER, c.wasm_file_name), filepath.Join(c.STATIC_FOLDER, "/app.wasm"))
	if err != nil {
		log.Println("BuildWASM error: ", err)
	}
}

func (c Compiler) buildWASM(input_go_file string, out_wasm_file string) error {

	var cmd *exec.Cmd

	// fmt.Println("WITH TINY GO?: ", c.with_tinyGo)
	// Ajustamos los parámetros de compilación según la configuración
	if c.with_tinyGo {
		// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
		cmd = exec.Command("tinygo", "build", "-o", out_wasm_file, "-target", "wasm", input_go_file)

	} else {
		// compilación normal...
		// fmt.Println("*** COMPILACIÓN WASM GO ***")
		cmd = exec.Command("go", "build", "-o", out_wasm_file, "-tags", "dev", "-ldflags", "-s -w", "-v", input_go_file)
		cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		// if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, string(output))
		return fmt.Errorf("al compilar a WebAssembly: %v", err)
	}

	// Verificamos si el archivo wasm se creó correctamente
	if _, err := os.Stat(out_wasm_file); err != nil {
		return fmt.Errorf("el archivo WebAssembly no se creó correctamente: %v", err)
	}

	// fmt.Printf("WebAssembly compilado correctamente y guardado en %s\n", out_wasm_file)

	return nil
}
