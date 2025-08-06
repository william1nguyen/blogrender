package post_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/william1nguyen/blogrender/internal/app/post"
)

func TestNewBlogPosts(t *testing.T) {
	const blogDir = "./posts"

	files, err := os.ReadDir(blogDir)
	if err != nil {
		log.Fatal(err)
	}

	cases := make([]struct {
		name     string
		filepath string
	}, len(files))

	for i, f := range files {
		cases[i] = struct {
			name     string
			filepath string
		}{
			name:     f.Name(),
			filepath: filepath.Join(blogDir, f.Name()),
		}
	}

	t.Run("it reads post markdown content", func(t *testing.T) {
		var posts []post.Post

		for _, c := range cases {
			p, err := post.NewPostFromFilePath(c.filepath)
			if err != nil {
				t.Fatal(err)
			}
			posts = append(posts, p)
		}

		approvals.VerifyJSONStruct(t, posts)
	})
}
