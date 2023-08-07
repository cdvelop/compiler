package compiler

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
)

// options ej:
// wasm:tinygo default:go
// menu:<code html here> default ""
// modules:<code html here> default ""
// icons:<code html svg sprite icon here> default ""
// project_dir:test default local run app dir
// modules_dir:c:\go\modules default HomeUserDir/Packages/go
// compile_dir:cmd default ""
// components_dir:c:\go\pkg default HomeUserDir/Packages/go
// theme_dir:c:\pkg\go\store default:HomeUserDir/Packages/go/platform
func Config(options ...string) *Compiler {

	c := Compiler{
		Page:                  model.Page{StyleSheet: "static/style.css", AppName: "apptest", AppVersion: "v0.0.0", UserName: "", UserArea: "", Message: "", Script: "static/main.js"},
		modules:               []*module{},
		components:            map[string]*component{},
		DirectoriesRegistered: map[string]struct{}{},
	}

	usr, err := user.Current()
	if err != nil {
		gotools.ShowErrorAndExit(err.Error())
	}
	c.components_dir = filepath.Join(usr.HomeDir, "Packages/go")

	current_dir, err := os.Getwd()
	if err != nil {
		gotools.ShowErrorAndExit(err.Error())
	}

	root_project_dir := current_dir
	gotools.RemoveSuffixIfPresent(&root_project_dir, "\\cmd")

	fmt.Println("DIRECTORIO ACTUAL: ", current_dir, " PROJECT ROOT: ", root_project_dir)

	c.DirectoriesRegistered[root_project_dir] = struct{}{}
	c.project_dir = current_dir

	c.modules_dir = c.components_dir
	c.theme_dir = filepath.Join(c.components_dir, "platform")

	var compile_dir string

	for _, option := range options {

		switch {

		case strings.Contains(option, "wasm:tinygo"):
			c.with_tinyGo = true

		case strings.Contains(option, "project_dir:"):
			gotools.ExtractTwoPointArgument(option, &c.project_dir)

		case strings.Contains(option, "compile_dir:"):
			gotools.ExtractTwoPointArgument(option, &compile_dir)

		case strings.Contains(option, "modules_dir:"):
			gotools.ExtractTwoPointArgument(option, &c.modules_dir)

		case strings.Contains(option, "components_dir:"):
			gotools.ExtractTwoPointArgument(option, &c.components_dir)

		case strings.Contains(option, "theme_dir:"):
			gotools.ExtractTwoPointArgument(option, &c.theme_dir)

		case strings.Contains(option, "menu:"):
			gotools.ExtractTwoPointArgument(option, &c.Page.Menu)

		case strings.Contains(option, "modules:"):
			gotools.ExtractTwoPointArgument(option, &c.Page.Modules)

		case strings.Contains(option, "icons:"):
			gotools.ExtractTwoPointArgument(option, &c.Page.SpriteIcons)

		}
	}

	c.wasm_file_name = "wasm_main.go"

	c.js_wasm_import = `const go = new Go();
	WebAssembly.instantiateStreaming(fetch("static/app.wasm"), go.importObject).then((result) => {
		go.run(result.instance);
	});`

	c.WORK_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend")
	c.BUILT_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend", "built")
	c.STATIC_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend", "built", "static")

	fmt.Println("THEME FOLDER: ", c.theme_dir)
	c.DirectoriesRegistered[c.theme_dir] = struct{}{}

	fmt.Println("PROJECT DIR: ", c.project_dir)
	fmt.Println("MODULES DIR: ", c.modules_dir)
	fmt.Println("COMPONENT DIR: ", c.components_dir)

	return &c
}

func (c *Compiler) CompileAllProject() {

	c.registrationFromCurrentDirectory()

	c.checkStaticFileFolders()

	c.copyStaticFilesFromUiTheme()

	c.webAssemblyCheck()

	c.rebuildAll()
}
