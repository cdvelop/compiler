package compiler

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/cdvelop/gotools"
	"github.com/tdewolff/minify"
	mincss "github.com/tdewolff/minify/css"
)

func (c Compiler) BuildCSS() {
	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	public_css := bytes.Buffer{}

	// fmt.Println(`1- comenzamos con el css del tema`)
	err := gotools.ReadFiles(c.theme_dir+"/css", ".css", &public_css)
	if err != nil {
		fmt.Println(err) // si hay error es por que no hay css en el tema
	}

	for _, c := range c.components {
		if c.folder_path != "" {
			gotools.ReadFiles(c.folder_path, ".css", &public_css)
		}
	}

	for _, m := range c.modules {
		c.attachInputsContentFromModule(m, ".css", &public_css)
	}
	// fmt.Println("4- >>> escribiendo archivos app.css y style.css")
	cssMinify(&public_css)

	gotools.FileWrite(c.STATIC_FOLDER+"/style.css", &public_css)

}

func cssMinify(data_in *bytes.Buffer) {

	m := minify.New()
	m.AddFunc("text/css", mincss.Minify)

	var temp_result bytes.Buffer
	err := m.Minify("text/css", &temp_result, data_in)

	if err != nil {
		log.Printf("Minification CSS error: %v\n", err)
		return
	}

	data_in.Reset()
	*data_in = temp_result

}
