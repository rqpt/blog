package renderer

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PostProp struct {
	Title string
	Body  template.HTML
}

var PostShow = template.Must(
	template.ParseFiles(
		filepath.Join("internal", "renderer", "templates", "posts", "show.html"),
		filepath.Join("internal", "renderer", "templates", "components", "layout.html"),
	),
)

func Render(
	w http.ResponseWriter,
	tmplType *template.Template,
	tmpl string,
	data any,
) {
	err := tmplType.ExecuteTemplate(w, fmt.Sprint(tmpl, ".html"), data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
