package compiler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cdvelop/fileserver"
	"github.com/tdewolff/minify"
	minjs "github.com/tdewolff/minify/js"
)

func (c *Compiler) BuildJS(event_name string) (err string) {
	const this = "BuildJS error "
	if event_name != "" {
		fmt.Println("Compilando JS..." + event_name)
	}

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	public_js := bytes.Buffer{}

	// fmt.Println(`0- agregamos js por defecto`)
	public_js.WriteString("'use strict';\n")

	// fmt.Println(`1- comenzamos con el js del tema`)
	fileserver.ReadFiles(filepath.Join(c.theme_dir, "js"), ".js", &public_js)

	// fmt.Println(`2- leer js publico de los componentes`)

	for _, comp := range c.components {

		if comp.folder_path != "" {

			fileserver.ReadFiles(filepath.Join(comp.folder_path, "js_global"), ".js", &public_js)
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

		fileserver.ReadFiles(filepath.Join(m.folder_path, "js_module"), ".js", &public_js)

		// fmt.Println(`agregamos js test si existiesen`)
		fileserver.ReadFiles(filepath.Join(m.folder_path, "js_test"), ".js", &public_js)

		// fmt.Println(`4- >>> escribiendo module JS: `, module.MainName)

		public_js.WriteString(moduleJsTemplate(m.name, funtions.String(), listener_add.String(), listener_rem.String()))

	}

	if c.wasm_build {
		var path_wasm_js string
		if c.with_tinyGo {
			path_wasm_js, err = wasmExecJsPathTinyGo()
		} else {
			path_wasm_js, err = wasmExecJsPathGo()
		}
		if err != "" {
			return this + err
		}

		// Leemos el contenido del archivo
		wasm_js_content, er := os.ReadFile(path_wasm_js)
		if er != nil {
			return this + er.Error()
		}

		public_js.Write(wasm_js_content)

		public_js.WriteString(c.js_wasm_import)
	}

	if c.minify {
		err = jsMinify(&public_js)
		if err != "" {
			return err
		}
	}

	return fileserver.FileWrite(filepath.Join(c.STATIC_FOLDER, "main.js"), &public_js)
}

func moduleJsTemplate(module_name, functions, listener_add, listener_rem string) string {

	if functions != "" || listener_add != "" && listener_rem != "" {

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

	return ""
}

func jsMinify(data_in *bytes.Buffer) (err string) {

	m := minify.New()
	m.AddFunc("text/javascript", minjs.Minify)

	var temp_result bytes.Buffer
	er := m.Minify("text/javascript", &temp_result, data_in)
	if er != nil {
		return "minification js error: " + er.Error()
	}

	data_in.Reset()
	*data_in = temp_result

	return ""
}
