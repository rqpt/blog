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
	"github.com/rqpt/blog/internal/slug"
)

type PostHandler struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	// TODO: Cater for both html and json requests
	renderer.Render(w, renderer.PostIndex, "layout", nil)
}

func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	wantsJSON := acceptsJSON(r)

	pathSlug := r.PathValue("post")
	pathSlugTitle, pathSlugID := slug.Parse(pathSlug)
	if pathSlugID == 0 {
		respondError(w, "Post not found", http.StatusNotFound, wantsJSON)
		return
	}

	post, err := h.queries.ShowPost(r.Context(), int64(pathSlugID))
	if err != nil {
		respondError(w, "Post not found", http.StatusNotFound, wantsJSON)
		return
	}

	postSlugTitle := slug.Make(post.Title)

	if postSlugTitle != pathSlugTitle && !wantsJSON {
		http.Redirect(
			w,
			r,
			fmt.Sprintf("/%v-%v", postSlugTitle, post.ID),
			http.StatusMovedPermanently,
		)
		return
	}

	if wantsJSON {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
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
