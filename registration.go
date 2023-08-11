package compiler

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/cdvelop/gomod"
	. "github.com/cdvelop/output"
)

func (c *Compiler) registrationFromCurrentDirectory() {

	current_project := filepath.Base(c.project_dir)

	if strings.Contains(current_project, "module") {
		fmt.Println("el proyecto actual es un modulo")

		new, err := c.createModule(current_project, c.project_dir)
		if err != nil {
			ShowErrorAndExit(err.Error())
		}

		c.addModule(new)

	} else {
		fmt.Println("=> directorio actual es un:")

		modules_names, components_names, err := gomod.GetSeparateUsedPackageNames(c.project_dir)
		if err != nil {
			ShowErrorAndExit(err.Error())
		}

		for _, module_name := range modules_names {
			fmt.Println("proyecto principal. creando módulo: ", module_name)

			module_dir := filepath.Join(c.modules_dir, module_name)

			new, err := c.createModule(module_name, module_dir)
			if err != nil {
				ShowErrorAndExit(err.Error())
			}

			c.addModule(new)

		}

		c.registerComponents(components_names)

	}

}
