package blogposts

import (
	"errors"
	"io/fs"
)

type Post struct {
	Title string
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
