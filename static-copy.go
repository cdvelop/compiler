package compiler

import (
	"os"
	"path/filepath"

	"github.com/cdvelop/fileserver"
)

func (c Compiler) copyStaticFilesFromUiTheme() (err string) {
	const this = "copyStaticFilesFromUiTheme error "
	// Definir las extensiones o tipos de archivo permitidos
	validExtensions := map[string]bool{".js": true, ".css": true, ".wasm": true}

	// Obtener la lista de archivos en la carpeta origen
	srcDir := filepath.Join(c.theme_dir, "/static")
	destDir := filepath.Join(c.BUILT_FOLDER, "/static")
	files, er := os.ReadDir(srcDir)
	if er != nil {
		return this + er.Error()
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
			if _, er := os.Stat(dest); os.IsNotExist(er) {
				// Si el archivo destino no existe, copiar el archivo
				err = fileserver.CopyFile(src, dest)
				if err != "" {
					return this + err
				}
			}
		}
	}

	return
}
