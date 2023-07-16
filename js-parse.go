package compiler

import (
	"bytes"
	"log"
	"text/template"
)

type parseJS struct {
	ModuleName string
	// FieldName  string
}

func parseModuleJS(p parseJS, data []byte) (string, error) {

	t, err := template.New("").Parse(string(data))
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}

	err = t.Execute(&buf, p)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return buf.String(), nil
}
