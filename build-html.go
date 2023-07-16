package compiler

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/cdvelop/gotools"
	"github.com/tdewolff/minify"
	minh "github.com/tdewolff/minify/html"
)

func (c *Compiler) BuildHTML() {
	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	template_html := c.makeHtmlTemplate()

	htmlMinify(&template_html)
	// crear archivo app html
	gotools.FileWrite(c.BUILT_FOLDER+"/index.html", &template_html)
}

func htmlMinify(data_in *bytes.Buffer) {

	m := minify.New()
	m.AddFunc("text/html", minh.Minify)

	var temp_result bytes.Buffer
	err := m.Minify("text/html", &temp_result, data_in)

	if err != nil {
		log.Printf("Minification HTML error: %v\n", err)
		return
	}

	data_in.Reset()
	*data_in = temp_result

}

func (c *Compiler) makeHtmlTemplate() (html bytes.Buffer) {

	data, err := os.ReadFile(c.theme_dir + "/index.html")
	if err != nil {
		fmt.Println("THEME FOLDER: ", c.theme_dir)
		fmt.Println("Error al leer el archivo:", err)
	}
	t, err := template.New("").Parse(string(data))
	if err != nil {
		log.Fatal(err)
		return
	}

	err = t.Execute(&html, c.Page)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
