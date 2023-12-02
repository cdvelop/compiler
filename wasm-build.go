package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (c Compiler) BuildWASM(event_name string) (err string) {
	const this = "BuildWASM error "
	if c.wasm_build {
		if event_name != "" {
			fmt.Println("Compilando WASM..." + event_name)
		}

		// fmt.Println("c.WORK_FOLDER:", c.WORK_FOLDER)

		// Ejecutar go mod tidy en el directorio del proyecto
		tidyCmd := exec.Command("go", "mod", "tidy")
		tidyCmd.Dir = c.project_dir
		tidyOutput, tidyErr := tidyCmd.CombinedOutput()
		if tidyErr != nil {
			return this + "al ejecutar 'go mod tidy': " + tidyErr.Error() + " " + string(tidyOutput)
		}

		var cmd *exec.Cmd

		// log.Println("*** c.test_wasm_folder: ", c.test_wasm_folder)

		input_go_file := filepath.Join(c.WORK_FOLDER, c.test_wasm_folder, c.wasm_file_name)

		out_wasm_file := filepath.Join(c.STATIC_FOLDER, "/app"+c.test_suffix+".wasm")

		// delete file anterior
		os.Remove(out_wasm_file)

		// log.Println("*** input_go_file: ", input_go_file)
		// Ajustamos los parámetros de compilación según la configuración
		if c.with_tinyGo {
			// fmt.Println("*** COMPILACIÓN WASM TINYGO ***")
			cmd = exec.Command("tinygo", "build", "-o", out_wasm_file, "-target", "wasm", input_go_file)

		} else {
			// compilación normal...
			cmd = exec.Command("go", "build", "-o", out_wasm_file, "-tags", "dev", "-ldflags", "-s -w", "-v", input_go_file)
			cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
		}

		output, er := cmd.CombinedOutput()
		if er != nil {
			return "error al compilar a WebAssembly error: " + er.Error() + " string(output):" + string(output)
		}

		// Verificamos si el archivo wasm se creó correctamente
		if _, er := os.Stat(out_wasm_file); er != nil {
			return "error el archivo WebAssembly no se creó correctamente: " + er.Error()
		}

		// fmt.Printf("WebAssembly compilado correctamente y guardado en %s\n", out_wasm_file)
	}

	return ""
}
