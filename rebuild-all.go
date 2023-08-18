package compiler

import (
	"fmt"
)

func (c Compiler) Rebuild() error {

	fmt.Println("... recompilando app archivos: html,css,js,wasm ...")

	err := c.BuildHTML("")
	if err != nil {
		return err
	}

	err = c.BuildJS("")
	if err != nil {
		return err
	}

	err = c.BuildCSS("")
	if err != nil {
		return err
	}

	err = c.BuildWASM("")
	if err != nil {
		return err
	}

	return nil
}
