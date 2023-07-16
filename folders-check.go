package compiler

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c Compiler) checkStaticFileFolders() {
	dirs := []string{filepath.Join(c.theme_dir, "/static"), filepath.Join(c.BUILT_FOLDER, "/static")}

	for _, dir := range dirs {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Printf("Error creando directorio %s: %v", dir, err)
				return
			}
			fmt.Printf("Directorio %s creado correctamente.\n", dir)
		} else if err != nil {
			fmt.Printf("Error al verificar directorio %s: %v", dir, err)
			return
		}
	}
}
