package compiler

import (
	"bytes"
	"fmt"
	"path/filepath"
	"time"

	"github.com/cdvelop/gotools"
	"github.com/tdewolff/minify"
	minjs "github.com/tdewolff/minify/js"
)

func (c *Compiler) BuildJS(event_name string) error {
	if event_name != "" {
		fmt.Println("Compilando JS..." + event_name)
	}

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	public_js := bytes.Buffer{}

	// fmt.Println(`0- agregamos js por defecto`)
	public_js.WriteString("'use strict';\n")

	// fmt.Println(`1- comenzamos con el js del tema`)
	gotools.ReadFiles(filepath.Join(c.theme_dir, "js"), ".js", &public_js)

	// fmt.Println(`2- leer js publico de los componentes`)

	for _, comp := range c.components {

		if comp.folder_path != "" {

			gotools.ReadFiles(filepath.Join(comp.folder_path, "js_global"), ".js", &public_js)
		}

	}

	// fmt.Println(`3- construir módulos js`)
	for _, m := range c.modules {
		funtions := bytes.Buffer{}
		listener_add := bytes.Buffer{}
		listener_rem := bytes.Buffer{}

		for _, component_name := range m.components_names {

			if comp, exist := c.components[component_name]; exist {

				attachJsToModuleFromFolder(comp, m.name, &funtions, &listener_add, &listener_rem)

			}
		}

		c.attachInputsContentFromModule(m, ".js", &public_js)

		gotools.ReadFiles(filepath.Join(m.folder_path, "js_module"), ".js", &public_js)

		// fmt.Println(`agregamos js test si existiesen`)
		gotools.ReadFiles(filepath.Join(m.folder_path, "js_test"), ".js", &public_js)

		// fmt.Println(`4- >>> escribiendo module JS: `, module.MainName)
		public_js.WriteString(moduleJsTemplate(m.name, funtions.String(), listener_add.String(), listener_rem.String()))

	}

	err := jsMinify(&public_js)
	if err != nil {
		return err
	}

	if c.wasm_build {
		if c.with_tinyGo {
			public_js.WriteString(addWasmJsTinyGo())
		} else {
			public_js.WriteString(addWasmJsGo())
		}
		public_js.WriteString(c.js_wasm_import)
	}

	err = gotools.FileWrite(filepath.Join(c.STATIC_FOLDER, "main.js"), &public_js)
	if err != nil {
		return err
	}

	return nil
}

func moduleJsTemplate(module_name, functions, listener_add, listener_rem string) string {
	return `MODULES['` + module_name + `'] = (function () {
	let crud = new Object();
	const module = document.getElementById('` + module_name + `');
	` + functions + `
	crud.ListenerModuleON = function () {
	 ` + listener_add + `
	};
	crud.ListenerModuleOFF = function () {
	 ` + listener_rem + `
	};
	return crud;
})();`
}

func jsMinify(data_in *bytes.Buffer) error {

	m := minify.New()
	m.AddFunc("text/javascript", minjs.Minify)

	var temp_result bytes.Buffer
	err := m.Minify("text/javascript", &temp_result, data_in)
	if err != nil {
		return fmt.Errorf("minification js error: %v", err)
	}

	data_in.Reset()
	*data_in = temp_result

	return nil
}
