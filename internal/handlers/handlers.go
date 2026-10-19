package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/rqpt/blog/internal/converter"
	"github.com/rqpt/blog/internal/db"
)

type Handler struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func (h *Handler) postIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<body>Welcome!</body>")
}

func (h *Handler) postCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	md := r.PostFormValue("body")
	html := converter.MdToHtml([]byte(md))

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

func NewHandler(conn *pgx.Conn) *Handler {
	return &Handler{
		conn:    conn,
		queries: db.New(conn),
	}
}
