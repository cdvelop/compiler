package module_info

import (
	"path/filepath"

	"github.com/cdvelop/model"
)

type module struct {
	*model.Module
}

func Get() *module {
	return &module{
		&model.Module{
			ModuleName: "module_info",
			Title:      "Información Plataforma TEST",
			IconID:     "icon-info",
			Areas:      map[string]string{"0": "", "t": ""},
		},
	}
}

func (m module) Icon(test_dir string) string {
	return filepath.Join(test_dir, "modules", m.ModuleName, m.IconID+".svg")
}
