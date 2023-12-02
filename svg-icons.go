package compiler

import (
	"github.com/cdvelop/fileserver"
)

func (c *Compiler) addSvgIcon(folder_path string) (err string) {
	// fmt.Println("AGREGAR ICONO SVG A HTML")
	var new_icons string

	err = fileserver.AddStringContendFromDirAndExtension(folder_path, ".svg", &new_icons)
	if err != "" {
		return "addSvgIcon error. NO SE LOGRO AGREGAR ICONO SVG DESDE " + folder_path
	}

	if _, exist := c.svg_icons[new_icons]; !exist {
		c.svg_icons[new_icons] = struct{}{}
	}
	return
}

func (c *Compiler) buildIconsSvg() {
	for icons := range c.svg_icons {
		c.Page.SpriteIcons += icons
	}
}
