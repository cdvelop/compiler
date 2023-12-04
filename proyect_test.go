package compiler_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cdvelop/compiler"
	"github.com/cdvelop/compiler/test/components/search"
	"github.com/cdvelop/compiler/test/modules/module_info"
	"github.com/cdvelop/compiler/test/modules/module_product"
	"github.com/cdvelop/fileserver"
	"github.com/cdvelop/gotools"
)

type key struct{}

func (key) LdFlagsEncryptionKey() map[string]string {
	return map[string]string{"main.key": "123"}
}

func Test_CompileProject(t *testing.T) {
	dir, _ := os.Getwd()
	test_dir := filepath.Join(dir, "test")

	// fmt.Println("***test_dir:", test_dir, "***")
	for _, wasm_compiler := range []string{"", "tinygo"} {

		for _, arg := range []string{"", "test:"} { // sin test: argumentos y con argumentos
			// var arg string

			c := compiler.Config(
				key{},
				"project_dir:"+filepath.Join(test_dir, "project"),
				"modules_dir:"+filepath.Join(test_dir, "modules"),
				"components_dir:"+filepath.Join(test_dir, "components"),
				wasm_compiler,
				arg,
			)
			// fmt.Println("***BUILT_FOLDER:", c.BUILT_FOLDER, "***")

			err := fileserver.DeleteFilesByExtension(c.BUILT_FOLDER, []string{".html"})
			if err != "" {
				t.Fatal(err)
				return
			}

			err = fileserver.DeleteFilesByExtension(c.STATIC_FOLDER, []string{".js", ".css", ".wasm"})
			if err != "" {
				t.Fatal(err)
				return
			}
			c.CompileAllProject()

			if arg == "test:" {
				arg = "_test"
			}

			err = fileserver.FindFilesWithNonZeroSize(c.BUILT_FOLDER, []string{"index.html", "style.css", "main.js", "app" + arg + ".wasm"})
			if err != "" {
				t.Fatal("Error argumento:", arg, err)
				return
			}

			resp := gotools.TextExists(filepath.Join(c.STATIC_FOLDER, "/style.css"), search.Css(test_dir))
			if resp == 0 {
				t.Fatal("EN style.css NO EXISTE: ", search.Css(test_dir))
				return
			}
			if resp > 1 {
				t.Fatal("EN style.css ESTA REPETIDO: ", search.Css(test_dir))
				return
			}

			if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Js(test_dir)) == 0 {
				t.Fatal("EN main.js NO EXISTE: ", search.Js(test_dir))
				return
			}

			if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Check().JsFunctionsExpected()) == 0 {
				t.Fatal("EN main.js NO EXISTE: ", search.Check().JsFunctionsExpected())
				return
			}

			if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.JsListener(test_dir)) == 0 {
				t.Fatal("EN main.js NO EXISTE: ", search.JsListener(test_dir))
				return
			}

			// removeEventListener se crea de forma dinámica
			if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Check().RemoveEventListenerExpected()) == 0 {
				t.Fatal("EN main.js NO EXISTE: ", search.Check().RemoveEventListenerExpected())
				return
			}

			//comprobar símbolos svg en html
			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_info.Get().Icon(test_dir)) == 0 {
				t.Fatal("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", module_info.Get().Icon(test_dir))
				return
			}

			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_info.Get().Icon(test_dir)) > 1 {
				t.Fatal("EN index.html icono repetido SÍMBOLO SVG ID : ", module_info.Get().Icon(test_dir))
				return
			}

			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_product.Get().Icon(test_dir)) == 0 {
				t.Fatal("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", module_product.Get().Icon(test_dir))
				return
			}

			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_product.Get().Icon(test_dir)) > 1 {
				t.Fatal("EN index.html icono repetido SÍMBOLO SVG ID : ", module_product.Get().Icon(test_dir))
				return
			}

			const icon_repeat = "icon-repeat"

			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", icon_repeat) == 0 {
				t.Fatal("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", icon_repeat)
				return
			}

			if gotools.TextExists(c.BUILT_FOLDER+"/index.html", icon_repeat) > 1 {
				t.Fatal("EN index.html icono repetido SÍMBOLO SVG ID : ", icon_repeat)
				return
			}

		}
	}

}
