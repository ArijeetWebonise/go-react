package templates

import (
	"bytes"
	"html/template"
)

type ITemplateParser interface {
	ParseTemplate([]string, interface{}) (string, error)
}

func (tp *TemplateParser) ParseTemplate(templateFileName []string, data interface{}) (string, error) {
	var parsedTemplate string
	t, err := template.ParseFiles(templateFileName...)
	if err != nil {
		return parsedTemplate, err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return parsedTemplate, err
	}
	parsedTemplate = buf.String()
	return parsedTemplate, nil
}

type TemplateParser struct {
}
