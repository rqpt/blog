package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rqpt/blog/internal/services"
)

var seedsPath = filepath.Join(
	"internal",
	"db",
	"seeds",
)

func Seed(ctx context.Context, db *pgxpool.Pool) error {
	queries := New(db)

	// When adding/removing seeds, increment/decrement the third arg.
	posts := make([]services.ConvertedPost, 0, 2)

	err := filepath.WalkDir(seedsPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		md, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		convertedPost, err := services.MdToHtml(string(md))
		if err != nil {
			return err
		}
		if convertedPost == nil {
			return errors.New("Markdown conversion encountered an issue.")
		}

		posts = append(posts, *convertedPost)

		return nil
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		_, err := queries.CreatePost(
			ctx,
			CreatePostParams{
				Title:       post.Title,
				Body:        post.Body,
				Frontmatter: post.Frontmatter,
			},
		)
		if err != nil {
			return fmt.Errorf("create post %q: %w", post.Title, err)
		}
	}

	return nil
}

func mustMarshal(v any) json.RawMessage {
	data, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("marshal seed data: %v", err))
	}

	return data
}
