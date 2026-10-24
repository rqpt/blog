package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rqpt/blog/internal/db"
	"github.com/rqpt/blog/internal/helpers"
	"github.com/rqpt/blog/internal/renderer"
	"github.com/rqpt/blog/internal/services"
)

func NewPostHandler(dbPool *pgxpool.Pool) *PostHandler {
	return &PostHandler{
		dbPool:  dbPool,
		queries: db.New(dbPool),
	}
}

type PostHandler struct {
	dbPool  *pgxpool.Pool
	queries *db.Queries
}

func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe set the Vary Accept header through middleware?
	w.Header().Add("Vary", "Accept")

	wantsJSON := helpers.AcceptsJSON(r)

	posts, err := h.queries.IndexPosts(r.Context())
	if err != nil {
		helpers.RespondError(w, "Posts not found", http.StatusNotFound, wantsJSON)
		return
	}

	if wantsJSON {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)
		return
	}

	renderer.Render(w, renderer.PostIndex, "layout", posts)
}

func (h *PostHandler) Show(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe set the Vary Accept header through middleware?
	w.Header().Add("Vary", "Accept")

	wantsJSON := helpers.AcceptsJSON(r)

	pathSlug := r.PathValue("post")

	pathID, err := helpers.ParseID(pathSlug)
	if pathID == 0 || err != nil {
		helpers.RespondError(w, "Post not found", http.StatusNotFound, wantsJSON)
		return
	}

	post, err := h.queries.ShowPost(r.Context(), pathID)
	if err != nil {
		helpers.RespondError(w, "Post not found", http.StatusNotFound, wantsJSON)
		return
	}

	if post.Slug != pathSlug && !wantsJSON {
		http.Redirect(
			w,
			r,
			fmt.Sprint("/", post.Slug),
			http.StatusMovedPermanently,
		)
		return
	}

	if wantsJSON {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
		return
	}

	data := renderer.PostShowProps{
		Title: post.Title,
		Body:  template.HTML(post.Body),
	}

	renderer.Render(w, renderer.PostShow, "layout", data)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: When we create a post, the markdown should be split into
	// frontmatter and content. Frontmatter should be saved on a new column
	// and used, among other things, for SEO meta tags. Content used for html.

	// TODO: Validation for unique title.
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
		helpers.RespondError(w, err.Error(), http.StatusInternalServerError, true)
		return
	}

	json.NewEncoder(w).Encode(newPost)
}
