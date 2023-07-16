package module_product

import (
	"path/filepath"

	"github.com/cdvelop/compiler/test/components/search"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

type module struct {
	*model.Module
}

func Get() *module {

	m := module{
		&model.Module{
			Theme:  nil,
			Name:   "module_product",
			Title:  "Productos TEST",
			IconID: "icon-products",
			UI:     module{},
			Areas:  []byte{'a', 't'},
		},
	}

	product_object := m.Object()
	product_object.AddModule(m.Module)

	search.Add().AddModule(m.Module)

	return &m
}

func (m module) Object() *model.Object {
	return &model.Object{
		Name:           "product",
		TextFieldNames: []string{},
		Fields: []model.Field{
			{Name: "name", Legend: "Nombre", Input: input.Text()},
			{Name: "mail", Legend: "Nombre", Input: input.Text()},
		},
	}
}

func (m module) Icon(test_dir string) string {
	return filepath.Join(test_dir, "modules", m.Name, m.IconID+".svg")
}

func (module) UserInterface(opt ...string) string {
	return "<h1>Modulo Productos</h1>"
}
