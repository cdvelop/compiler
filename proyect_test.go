package compiler_test

import (
	"log"
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

func Test_CompileProject(t *testing.T) {
	dir, _ := os.Getwd()
	test_dir := filepath.Join(dir, "test")

	for _, wasm_compiler := range []string{"", "tinygo"} {

		c := compiler.Config(
			"project_dir:"+filepath.Join(test_dir, "project"),
			"modules_dir:"+filepath.Join(test_dir, "modules"),
			"components_dir:"+filepath.Join(test_dir, "components"),
			wasm_compiler,
		)

		err := fileserver.DeleteFilesByExtension(c.BUILT_FOLDER, []string{".html"})
		if err != "" {
			log.Fatalln(err)
		}

		err = fileserver.DeleteFilesByExtension(c.STATIC_FOLDER, []string{".js", ".css", ".wasm"})
		if err != "" {
			log.Fatalln(err)
		}
		c.CompileAllProject()

		err = fileserver.FindFilesWithNonZeroSize(c.BUILT_FOLDER, []string{"index.html", "style.css", "main.js", "app.wasm"})
		if err != "" {
			log.Fatal("Error:", err)
		}

		resp := gotools.TextExists(filepath.Join(c.STATIC_FOLDER, "/style.css"), search.Css(test_dir))
		if resp == 0 {
			log.Fatalln("EN style.css NO EXISTE: ", search.Css(test_dir))
		}
		if resp > 1 {
			log.Fatalln("EN style.css ESTA REPETIDO: ", search.Css(test_dir))
		}

		if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Js(test_dir)) == 0 {
			log.Fatalln("EN main.js NO EXISTE: ", search.Js(test_dir))
		}

		if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Check().JsFunctionsExpected()) == 0 {
			log.Fatalln("EN main.js NO EXISTE: ", search.Check().JsFunctionsExpected())
		}

		if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.JsListener(test_dir)) == 0 {
			log.Fatalln("EN main.js NO EXISTE: ", search.JsListener(test_dir))
		}

		// removeEventListener se crea de forma dinámica
		if gotools.TextExists(c.STATIC_FOLDER+"/main.js", search.Check().RemoveEventListenerExpected()) == 0 {
			log.Fatalln("EN main.js NO EXISTE: ", search.Check().RemoveEventListenerExpected())
		}

		//comprobar símbolos svg en html
		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_info.Get().Icon(test_dir)) == 0 {
			log.Fatalln("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", module_info.Get().Icon(test_dir))
		}

		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_info.Get().Icon(test_dir)) > 1 {
			log.Fatalln("EN index.html icono repetido SÍMBOLO SVG ID : ", module_info.Get().Icon(test_dir))
		}

		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_product.Get().Icon(test_dir)) == 0 {
			log.Fatalln("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", module_product.Get().Icon(test_dir))
		}

		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", module_product.Get().Icon(test_dir)) > 1 {
			log.Fatalln("EN index.html icono repetido SÍMBOLO SVG ID : ", module_product.Get().Icon(test_dir))
		}

		const icon_repeat = "icon-repeat"

		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", icon_repeat) == 0 {
			log.Fatalln("EN index.html NO SE CREO EL SÍMBOLO SVG ID : ", icon_repeat)
		}

		if gotools.TextExists(c.BUILT_FOLDER+"/index.html", icon_repeat) > 1 {
			log.Fatalln("EN index.html icono repetido SÍMBOLO SVG ID : ", icon_repeat)
		}

	}

}
