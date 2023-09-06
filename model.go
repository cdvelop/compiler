package compiler

import "github.com/cdvelop/model"

type Compiler struct {
	project_dir    string // directorio actual
	modules_dir    string // directorio módulos
	components_dir string // directorio paquetes
	theme_dir      string // directorio tema default components_dir + \platform

	model.Page

	minify bool

	wasm_build  bool
	with_tinyGo bool

	//módulos registrados
	modules []*module

	//componentes registrados
	components map[string]*component

	//id icono mas contenido
	svg_icons map[string]struct{}

	wasm_file_name string
	js_wasm_import string

	WORK_FOLDER   string
	BUILT_FOLDER  string
	STATIC_FOLDER string

	DirectoriesRegistered map[string]struct{}

	log bool
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
