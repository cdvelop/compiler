package compiler

import (
	"bytes"
	"log"
	"os"
	"strings"
)

// path ej: "modules/users/js_module","ui/components/form/js_module"
func attachJsToModuleFromFolder(comp *component, module_name string, funtions, listener_add, listener_rem *bytes.Buffer) {

	// adjuntar js desde carpeta
	if comp.folder_path != "" {
		path := comp.folder_path + "/js_module"

		files, err := os.ReadDir(path)
		if err == nil {
			// fmt.Printf("directorio %v de %v no encontrado\n", path, module.MainName)

			for _, file := range files {

				data, err := os.ReadFile(path + "/" + file.Name())
				if err != nil {
					log.Fatalf("error: archivo %v/%v no existe %v", path, file.Name(), err)
				}

				parsed_js, _ := parseModuleJS(parseJS{
					ModuleName: module_name,
				}, data)

				// fmt.Println("PARSE: ", parsed_js.String())
				// fmt.Println("FILE NAME: ", file.Name())

				if strings.Contains(parsed_js, "addEventListener") {

					listener_add.WriteString(parsed_js + "\n")

					// reemplazar todas las ocurrencias de "addEventListener" por "removeEventListener"
					rem_listeners := strings.ReplaceAll(parsed_js, "addEventListener", "removeEventListener")

					listener_rem.WriteString(rem_listeners + "\n")
				} else {

					funtions.WriteString(parsed_js + "\n")
				}

			}

		}

	}

}
