package renderer

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PostProps struct {
	Title string
	Body  template.HTML
}

var templatesPath = filepath.Join(
	"internal",
	"renderer",
	"resources",
	"templates",
)

var PostShow = template.Must(
	template.ParseFiles(
		filepath.Join(templatesPath, "components", "layout.html"),
		filepath.Join(templatesPath, "posts", "show.html"),
	),
)

var PostIndex = template.Must(
	template.ParseFiles(
		filepath.Join(templatesPath, "components", "layout.html"),
		filepath.Join(templatesPath, "posts", "index.html"),
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
