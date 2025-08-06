package post

import (
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func (p *Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
