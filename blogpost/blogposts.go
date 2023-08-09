package blogposts

import (
	"errors"
	"io/fs"
)

type Posts struct {
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func NewPostsFromFS(filesystem fs.FS) ([]Posts, error) {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Posts
	for range dir {
		posts = append(posts, Posts{})
	}
	return posts, nil
}
