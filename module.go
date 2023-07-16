package compiler

import (
	"fmt"

	"github.com/cdvelop/gomod"
	"github.com/cdvelop/gotools"
)

func (c *Compiler) createModule(name, path string, components_names ...string) (m *module, err error) {

	if len(components_names) == 0 {
		_, components_names, err = gomod.GetSeparateUsedPackageNames(path)
		if err != nil {
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

		fmt.Println("AGREGAR ICONO SVG A HTML")

		err := gotools.AddStringContendFromDirAndExtension(new.folder_path, ".svg", &c.Page.SpriteIcons)
		if err != nil {
			gotools.ShowErrorAndExit(err.Error())
		}

	}
}
