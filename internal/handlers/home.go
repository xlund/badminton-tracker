package handlers

import (
	"net/http"
	"text/template"

	"github.com/xlund/badminton-tracker/web/templates"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("index").Parse(templates.IndexHTML)
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, nil)
	}
}
