package compiler

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
	. "github.com/cdvelop/output"
)

// options ej:
// no_minify
// tinygo (wasm compiler) default go
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
		minify:                true,
	}

	usr, err := user.Current()
	if err != nil {
		ShowErrorAndExit(err.Error())
	}
	c.components_dir = filepath.Join(usr.HomeDir, "Packages/go")

	current_dir, err := os.Getwd()
	if err != nil {
		ShowErrorAndExit(err.Error())
	}

	root_project_dir := current_dir
	gotools.RemoveSuffixIfPresent(&root_project_dir, "\\cmd")

	fmt.Println("DIRECTORIO ACTUAL: ", current_dir, " PROJECT ROOT: ", root_project_dir)

	c.DirectoriesRegistered[root_project_dir] = struct{}{}
	c.project_dir = current_dir

	c.modules_dir = c.components_dir
	c.theme_dir = filepath.Join(c.components_dir, "platform")

	var compile_dir string

	for _, arg := range os.Args {
		switch {

		case arg == "no_minify":
			c.minify = false

		case arg == "tinygo":
			c.with_tinyGo = true

		case strings.Contains(arg, "theme_dir:"):
			gotools.ExtractTwoPointArgument(arg, &c.theme_dir)
		}
	}

	for _, option := range options {

		switch {

		case option == "tinygo":
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

	c.WORK_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend")
	c.BUILT_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend", "built")
	c.STATIC_FOLDER = filepath.Join(c.project_dir, compile_dir, "frontend", "built", "static")

	PrintInfo("THEME FOLDER: " + c.theme_dir)
	c.DirectoriesRegistered[c.theme_dir] = struct{}{}

	PrintInfo("PROJECT DIR: " + c.project_dir)
	PrintInfo("MODULES DIR: " + c.modules_dir)
	PrintInfo("COMPONENT DIR: " + c.components_dir)

	return &c
}

func (c *Compiler) CompileAllProject() {

	c.registrationFromCurrentDirectory()

	c.checkStaticFileFolders()

	c.copyStaticFilesFromUiTheme()

	c.webAssemblyCheck()

	err := c.Rebuild()
	if err != nil {
		ShowErrorAndExit(err)
	}
}

func (c *Compiler) ThemeDir() string {
	return c.theme_dir
}
