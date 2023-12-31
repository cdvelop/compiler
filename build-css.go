package compiler

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/cdvelop/fileserver"
	"github.com/tdewolff/minify"
	mincss "github.com/tdewolff/minify/css"
)

func (c Compiler) BuildCSS(event_name string) (err string) {
	const this = "BuildCSS error "
	if event_name != "" {
		fmt.Println("Compilando CSS..." + event_name)
	}

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	public_css := bytes.Buffer{}

	// fmt.Println(`1- comenzamos con el css del tema`)
	err = fileserver.ReadFiles(filepath.Join(c.theme_dir, "css"), ".css", &public_css)
	if err != "" {
		return this + "el tema no contiene la carpeta /css"
	}

	for _, c := range c.components {
		if c.folder_path != "" {
			fileserver.ReadFiles(c.folder_path, ".css", &public_css)
		}
	}

	for _, m := range c.modules {
		c.attachInputsContentFromModule(m, ".css", &public_css)
	}

	// fmt.Println("4- >>> escribiendo archivos app.css y style.css")
	if c.minify {
		cssMinify(&public_css)
	}

	fileserver.FileWrite(filepath.Join(c.STATIC_FOLDER, "style.css"), &public_css)

	return ""
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
