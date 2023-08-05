package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Posts struct {
}

func NewPostsFromFS(filesystem fstest.MapFS) []Posts {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Posts
	for range dir {
		posts = append(posts, Posts{})
	}
	return posts
}
