package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/rqpt/blog/internal/db"
	"github.com/rqpt/blog/internal/renderer"
	"github.com/rqpt/blog/internal/services"
)

type PostHandler struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	renderer.Render(w, renderer.PostIndex, "layout", nil)
}

func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	// TODO: Make sure we can serve as either html or json.
	pathSlug := r.PathValue("post")
	pathSlugTitle, pathSlugID := getTitleAndIDFromSlug(pathSlug)
	if pathSlugID == 0 {
		http.NotFound(w, r)
		return
	}

	post, err := h.queries.ShowPost(r.Context(), int64(pathSlugID))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	postSlugTitle, err := makeSlug(post.Title)
	if err != nil {
		http.Error(
			w,
			"Post Title could not be converted to a slug",
			http.StatusInternalServerError,
		)
		return
	}

	if postSlugTitle != pathSlugTitle {
		http.Redirect(
			w,
			r,
			fmt.Sprintf("/%v-%v", postSlugTitle, post.ID),
			http.StatusMovedPermanently,
		)
		return
	}

	data := renderer.PostProp{
		Title: post.Title,
		Body:  template.HTML(post.Body),
	}

	renderer.Render(w, renderer.PostShow, "layout", data)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	md := r.PostFormValue("body")
	html := services.MdToHtml([]byte(md))

	newPost, err := h.queries.CreatePost(
		r.Context(),
		db.CreatePostParams{
			Title: r.PostFormValue("title"),
			Body:  string(html),
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(newPost)
}

func NewPostHandler(conn *pgx.Conn) *PostHandler {
	return &PostHandler{
		conn:    conn,
		queries: db.New(conn),
	}
}
