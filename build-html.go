package compiler

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/cdvelop/gotools"
	"github.com/tdewolff/minify"
	minh "github.com/tdewolff/minify/html"
)

func (c *Compiler) BuildHTML(event_name string) error {
	if event_name != "" {
		fmt.Println("Compilando HTML..." + event_name)

	}

	c.buildIconsSvg()

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	template_html, err := c.makeHtmlTemplate()
	if err != nil {
		return err
	}

	if c.minify {
		err = htmlMinify(&template_html)
		if err != nil {
			return err
		}
	}

	// crear archivo app html
	gotools.FileWrite(filepath.Join(c.BUILT_FOLDER, "index.html"), &template_html)

	return nil
}

func htmlMinify(data_in *bytes.Buffer) error {

	m := minify.New()
	m.AddFunc("text/html", minh.Minify)

	var temp_result bytes.Buffer
	err := m.Minify("text/html", &temp_result, data_in)
	if err != nil {
		return fmt.Errorf("minification html error: %v", err)
	}

	data_in.Reset()
	*data_in = temp_result

	return nil
}

func (c *Compiler) makeHtmlTemplate() (html bytes.Buffer, err error) {

	data, err := os.ReadFile(filepath.Join(c.theme_dir, "index.html"))
	if err != nil {
		fmt.Println("THEME FOLDER: ", c.theme_dir)
		fmt.Println("Error al leer el archivo:", err)
	}
	t, err := template.New("").Parse(string(data))
	if err != nil {
		log.Fatal(err)
		return
	}

	// template.HTMLEscapeString()
	if c.Page.JsonBootActions == "" {
		c.Page.JsonBootActions = `{{.JsonBootActions}}`
		c.Page.JsonBootTests = `{{.JsonBootTests}}`
	}

	err = t.Execute(&html, c.Page)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}
