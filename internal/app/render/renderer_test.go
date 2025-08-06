package render_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/william1nguyen/blogrender/internal/app/post"
	"github.com/william1nguyen/blogrender/internal/app/render"
)

func TestRender(t *testing.T) {
	var (
		p = post.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}

		posts = []post.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}
	)

	postRenderer, err := render.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, p); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		post = post.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	for b.Loop() {
		postRenderer, err := render.NewPostRenderer()

		if err != nil {
			b.Fatal(err)
		}

		postRenderer.Render(io.Discard, post)
	}
}
