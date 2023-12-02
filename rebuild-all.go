package compiler

import (
	"fmt"
)

func (c Compiler) Rebuild() (err string) {

	fmt.Println("... recompilando app archivos: html,css,js,wasm ...")

	err = c.BuildHTML("")
	if err != "" {
		return
	}

	err = c.BuildJS("")
	if err != "" {
		return
	}

	err = c.BuildCSS("")
	if err != "" {
		return
	}

	err = c.BuildWASM("")
	if err != "" {
		return
	}

	return ""
}
