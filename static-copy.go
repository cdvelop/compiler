package compiler

import (
	"os"
	"path/filepath"

	"github.com/cdvelop/gotools"
)

func (c Compiler) copyStaticFilesFromUiTheme() {
	// Definir las extensiones o tipos de archivo permitidos
	validExtensions := map[string]bool{".js": true, ".css": true, ".wasm": true}

	// Obtener la lista de archivos en la carpeta origen
	srcDir := filepath.Join(c.theme_dir, "/static")
	destDir := filepath.Join(c.BUILT_FOLDER, "/static")
	files, err := os.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}

	// Recorrer la lista de archivos
	for _, file := range files {
		// Verificar si el archivo no es de una extensión prohibida
		ext := filepath.Ext(file.Name())
		if !validExtensions[ext] {
			// Obtener la ruta completa del archivo origen y destino
			src := filepath.Join(srcDir, file.Name())
			dest := filepath.Join(destDir, file.Name())

			// Verificar si el archivo destino ya existe
			if _, err := os.Stat(dest); os.IsNotExist(err) {
				// Si el archivo destino no existe, copiar el archivo
				err := gotools.CopyFile(src, dest)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
