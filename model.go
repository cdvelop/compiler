package compiler

import "github.com/cdvelop/model"

type Compiler struct {
	project_dir    string // directorio actual
	modules_dir    string // directorio módulos
	components_dir string // directorio paquetes
	theme_dir      string // directorio tema default components_dir + \platform

	model.Page

	wasm_build  bool
	with_tinyGo bool

	//módulos registrados
	modules []*module

	//componentes registrados
	components map[string]*component

	WORK_FOLDER   string
	BUILT_FOLDER  string
	STATIC_FOLDER string
}

type module struct {
	name             string
	components_names []string
	folder_path      string
}

type component struct {
	name        string
	folder_path string
}
