package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rqpt/blog/internal/services"
)

func Seed(ctx context.Context, db *pgxpool.Pool) error {
	queries := New(db)

	posts := []services.ConvertedPost{
		{
			Title: "Test",
			Body:  "<h1>Hello, world!</h1>",
			Frontmatter: mustMarshal(
				services.Frontmatter{
					Title:       "test",
					Description: "test",
					Tags:        []string{"test1", "test2"},
				},
			),
		},
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
