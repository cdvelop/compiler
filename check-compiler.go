package compiler

import (
	"fmt"

	"github.com/cdvelop/gotools"
)

func (c Compiler) compilerCheck() {

	c.BuildHTML()

	err := gotools.FindFilesWithNonZeroSize(c.BUILT_FOLDER, []string{"style.css", "main.js"})
	if err != nil {
		fmt.Println(err, "... recompilando proyecto archivos: html,css,js,wasm ...")

		c.BuildJS()

		c.BuildCSS()

		c.BuildWASM()

	}
}
