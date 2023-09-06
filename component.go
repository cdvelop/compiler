package compiler

import (
	"os"
	"path/filepath"
)

func (c *Compiler) registerComponents(components_names []string) {

	for _, component_name := range components_names {

		if _, no_exist := c.components[component_name]; !no_exist {

			folder_path := filepath.Join(c.components_dir, component_name)
			_, err := os.Stat(folder_path)
			if err == nil {
				// fmt.Println("directorio no existe nada que eliminar")
				// fmt.Println("*** ERROR COMPONENTE:", component_name, err)
				// fmt.Println("*** ERROR PATH NO EXISTE:", folder_path)

				// fmt.Println("*** OK AGREGANDO COMPONENTE:", component_name)
				// fmt.Println("*** PATH:", folder_path)

				new_component := component{
					name:        component_name,
					folder_path: folder_path,
				}
				c.DirectoriesRegistered[new_component.folder_path] = struct{}{}

				c.components[component_name] = &new_component

				c.addSvgIcon(folder_path)
			}
		}

	}
}
