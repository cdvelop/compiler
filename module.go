package compiler

import (
	"fmt"

	"github.com/cdvelop/gomod"
)

func (c *Compiler) createModule(name, path string, components_names ...string) (m *module, err string) {

	if len(components_names) == 0 {
		_, components_names, err = gomod.GetSeparateUsedPackageNames(path)
		if err != "" {
			return nil, err
		}
	}

	m = &module{
		name:             name,
		components_names: components_names,
		folder_path:      path,
	}

	return
}

func (c *Compiler) addModule(new *module) {
	var module_found *module

	for _, m := range c.modules {
		if m != nil && m.name == new.name {
			module_found = new
		}
	}

	if module_found == nil && new != nil {
		c.modules = append(c.modules, new)

		c.registerComponents(new.components_names)

		fmt.Println("modulo: ", new.name, " componentes: ", new.components_names)

		c.addSvgIcon(new.folder_path)

		c.DirectoriesRegistered[new.folder_path] = struct{}{}

	}
}
