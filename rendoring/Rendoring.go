package rendoring

import (
	"html/template"
	"net/http"
)

type Page struct {
	Url string
	Error string
	SubmissionNotification string
	Words []string
}

var templates = template.Must(template.ParseFiles("view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	var err error
	if p != nil {
		err = templates.ExecuteTemplate(w, tmpl+".html", p)
	} else {
		err = templates.ExecuteTemplate(w, tmpl+".html", nil)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
