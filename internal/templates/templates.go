package templates

import "html/template"

var Template *template.Template

func Read() error {
	tmpl, err := template.ParseGlob("./web/templates/*.html")
	Template = tmpl

	return err
}
