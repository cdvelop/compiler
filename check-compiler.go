package compiler

import (
	"fmt"
)

func (c Compiler) rebuildAll() {

	fmt.Println("... recompilando proyecto archivos: html,css,js,wasm ...")

	c.BuildHTML()

	c.BuildJS()

	c.BuildCSS()

	c.BuildWASM()

}
