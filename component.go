package compiler

import "path/filepath"

func (c *Compiler) registerComponents(components_names []string) {

	for _, component_name := range components_names {

		if _, no_exist := c.components[component_name]; !no_exist {

			new_component := component{
				name:        component_name,
				folder_path: filepath.Join(c.components_dir, component_name),
			}
			c.DirectoriesRegistered[new_component.folder_path] = struct{}{}

			c.components[component_name] = &new_component
		}

	}
}
