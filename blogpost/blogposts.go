package blogposts

import (
	"io/fs"
)

type Posts struct {
}

func NewPostsFromFS(filesystem fs.FS) ([]Posts, error) {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Posts
	for range dir {
		posts = append(posts, Posts{})
	}
	return posts, nil
}
