package search

import (
	"path/filepath"

	"github.com/cdvelop/model"
)

type search struct{}

func Add() *model.Object {

	return &model.Object{
		Name: "search",
	}
}

func Check() search {
	return search{}
}

func (s search) JsFunctionsExpected() string {
	return "console.log('función componente search modulo: module_product')"
}

// esta función es solo para comparar en el test ya que se crea de forma dinámica
func (search) RemoveEventListenerExpected() string {
	return "btn.removeEventListener('click',MySearchFunction);"
}

func Css(test_dir string) string {
	return filepath.Join(test_dir, "/components/search/style.css")
}

func Js(test_dir string) string {
	return filepath.Join(test_dir, "components/search/js_global/main.js")
}

func JsListener(test_dir string) string {
	return filepath.Join(test_dir, "components/search/js_module/listener-add.js")

}
