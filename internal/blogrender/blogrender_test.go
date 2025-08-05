package blogrender_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/william1nguyen/blogrender/internal/blogrender"
)

func TestRender(t *testing.T) {
	var (
		post = blogrender.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	buf := bytes.Buffer{}
	err := blogrender.Render(&buf, post)

	if err != nil {
		t.Fatal(err)
	}
	t.Parallel()
	approvals.VerifyString(t, buf.String())
}

func BenchmarkRender(b *testing.B) {
	var (
		post = blogrender.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	for b.Loop() {
		blogrender.Render(io.Discard, post)
	}
}
