package compiler

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/model"
	. "github.com/cdvelop/output"
)

// extension: .js, .css
func (c Compiler) attachInputsContentFromModule(m *module, extension string, out *bytes.Buffer) {

	if m.folder_path != "" && extension != "" {
		// obtenemos los nombres de  los input de tipo go usados del modulo
		input_names, err := fileserver.GetNamesFromDirectoryExtensionAndPattern(m.folder_path, ".go", model.INPUT_PATTERN)
		if err != "" {
			ShowErrorAndExit(err)
		}

		// OBTENER UBICACIÓN POR DEFECTO INPUTS
		inputs_path := filepath.Join(c.components_dir, "input")

		// fmt.Println("INPUT PATH: ", inputs_path)

		for _, input_name := range input_names {

			file_path := filepath.Join(inputs_path, input_name+extension)

			// fmt.Println("FILE PATH INPUT: ", file_path)

			content, err := os.ReadFile(file_path)
			if err == nil {
				out.Write(content)
			}

		}

	}

}
