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
		queries: db.New(dbPool),
	}
}

type PostHandler struct {
	queries *db.Queries
}

func (h *PostHandler) Index(w http.ResponseWriter, r *http.Request) {
	wantsJSON := helpers.AcceptsJSON(r)

	posts, err := h.queries.IndexPosts(r.Context())
	if err != nil {
		helpers.RespondError(
			w,
			"Posts not found",
			http.StatusNotFound,
			wantsJSON,
		)
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
	wantsJSON := helpers.AcceptsJSON(r)

	pathID, pathSlug, ok := h.parsePathSlug(w, r, wantsJSON)
	if !ok {
		return
	}

	post, err := h.queries.ShowPost(r.Context(), pathID)
	if err != nil {
		helpers.RespondError(
			w,
			"Post not found",
			http.StatusNotFound,
			wantsJSON,
		)
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
	// TODO: Validation for unique title.
	// TODO: Validation that md form value is present.
	convertedPost, ok := h.parseMarkdown(w, r, true)
	if !ok {
		return
	}

	newPost, err := h.queries.CreatePost(
		r.Context(),
		db.CreatePostParams{
			Title:       convertedPost.Title,
			Body:        convertedPost.Body,
			Frontmatter: convertedPost.Frontmatter,
		},
	)
	if err != nil {
		helpers.RespondError(
			w,
			err.Error(),
			http.StatusInternalServerError,
			true,
		)
		return
	}

	json.NewEncoder(w).Encode(newPost)
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Validation for unique title.
	pathID, _, ok := h.parsePathSlug(w, r, true)
	if !ok {
		return
	}

	convertedPost, ok := h.parseMarkdown(w, r, true)
	if !ok {
		return
	}

	err := h.queries.UpdatePost(
		r.Context(),
		db.UpdatePostParams{
			ID:          pathID,
			Title:       convertedPost.Title,
			Body:        convertedPost.Body,
			Frontmatter: convertedPost.Frontmatter,
		},
	)
	if err != nil {
		helpers.RespondError(
			w,
			err.Error(),
			http.StatusInternalServerError,
			true,
		)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Successfully updated post #%d", pathID),
	})
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	pathID, _, ok := h.parsePathSlug(w, r, true)
	if !ok {
		return
	}

	err := h.queries.DeletePost(r.Context(), pathID)
	if err != nil {
		helpers.RespondError(
			w,
			err.Error(),
			http.StatusInternalServerError,
			true,
		)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Successfully deleted post #%d", pathID),
	})
}

func (h *PostHandler) parseMarkdown(
	w http.ResponseWriter,
	r *http.Request,
	wantsJson bool,
) (*services.ConvertedPost, bool) {
	md := r.PostFormValue("md")

	convertedPost, err := services.MdToHtml(md)
	if err != nil {
		helpers.RespondError(
			w,
			err.Error(),
			http.StatusBadRequest,
			wantsJson,
		)
		return nil, false
	}
	if convertedPost == nil {
		helpers.RespondError(
			w,
			"failed to convert Markdown",
			http.StatusInternalServerError,
			wantsJson,
		)
		return nil, false
	}

	return convertedPost, true
}

func (h *PostHandler) parsePathSlug(
	w http.ResponseWriter,
	r *http.Request,
	wantsJson bool,
) (int64, string, bool) {
	pathSlug := r.PathValue("post")

	pathID, err := helpers.ParseID(pathSlug)
	if pathID == 0 || err != nil {
		helpers.RespondError(
			w,
			"Post not found",
			http.StatusNotFound,
			wantsJson,
		)
		return 0, "", false
	}

	return pathID, pathSlug, true
}
