package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (c Compiler) BuildWASM(event_name string) error {

	if c.wasm_build {
		if event_name != "" {
			fmt.Println("Compilando WASM..." + event_name)
		}

		var cmd *exec.Cmd

		input_go_file := filepath.Join(c.WORK_FOLDER, c.wasm_file_name)

		out_wasm_file := filepath.Join(c.STATIC_FOLDER, "/app.wasm")

		// delete file anterior
		os.Remove(out_wasm_file)

		// fmt.Println("WITH TINY GO?: ", c.with_tinyGo)
		// Ajustamos los parámetros de compilación según la configuración
		if c.with_tinyGo {
			// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
			cmd = exec.Command("tinygo", "build", "-o", out_wasm_file, "-target", "wasm", input_go_file)

		} else {
			// compilación normal...
			cmd = exec.Command("go", "build", "-o", out_wasm_file, "-tags", "dev", "-ldflags", "-s -w", "-v", input_go_file)
			cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error al compilar a WebAssembly: %v %v", err, string(output))
		}

		// Verificamos si el archivo wasm se creó correctamente
		if _, err := os.Stat(out_wasm_file); err != nil {
			return fmt.Errorf("error el archivo WebAssembly no se creó correctamente: %v", err)
		}

		// fmt.Printf("WebAssembly compilado correctamente y guardado en %s\n", out_wasm_file)
	}

	return nil
}
