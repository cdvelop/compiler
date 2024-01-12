package compiler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/cdvelop/fileserver"
	"github.com/tdewolff/minify"
	minh "github.com/tdewolff/minify/html"
)

func (c *Compiler) BuildHTML(event_name string) (err string) {
	const this = "BuildHTML error "
	if event_name != "" {
		fmt.Println("Compilando HTML..." + event_name)

	}

	c.buildIconsSvg()

	time.Sleep(10 * time.Millisecond) // Esperar antes de intentar leer el archivo de nuevo

	template_html, err := c.makeHtmlTemplate()
	if err != "" {
		return this + err
	}

	if c.minify {
		err = htmlMinify(&template_html)
		if err != "" {
			return this + err
		}
	}

	// crear archivo app html

	return fileserver.FileWrite(filepath.Join(c.BUILT_FOLDER, "index.html"), &template_html)
}

func htmlMinify(data_in *bytes.Buffer) (err string) {

	m := minify.New()
	m.AddFunc("text/html", minh.Minify)

	var temp_result bytes.Buffer
	er := m.Minify("text/html", &temp_result, data_in)
	if er != nil {
		return "htmlMinify error: " + err
	}

	data_in.Reset()
	*data_in = temp_result

	return ""
}

func (c *Compiler) makeHtmlTemplate() (html bytes.Buffer, err string) {
	const this = "makeHtmlTemplate error "
	data, er := os.ReadFile(filepath.Join(c.theme_dir, "index.html"))
	if er != nil {
		err = this + "THEME FOLDER: " + c.theme_dir + " " + er.Error()
		return
	}
	t, er := template.New("").Parse(string(data))
	if er != nil {
		err = this + er.Error()
		return
	}

	// template.HTMLEscapeString()
	if c.Page.JsonBootActions == "" {
		c.Page.JsonBootActions = `{{.JsonBootActions}}`
	}

	c.Page.AppName = c.Config.AppName()
	c.Page.AppVersion = c.Config.AppVersion()
	c.Page.StyleSheet = "static/style.css" + c.versionStatics()
	c.Page.Script = "static/main.js" + c.versionStatics()

	er = t.Execute(&html, c.Page)
	if er != nil {
		err = this + er.Error()
	}

	return
}
