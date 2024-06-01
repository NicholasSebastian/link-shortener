package utils

import (
	"net/http"
	"path/filepath"
)

// Overrides the default filesystem behavior to make routes fallback to '.html'.
// I don't know why this isn't the default behaviour for a HTTP static file server.

type HtmlDir struct {
	Dir *http.Dir
}

func (fs HtmlDir) Open(name string) (http.File, error) {
	if name != "/" && filepath.Ext(name) == "" {
		name += ".html"
	}
	return fs.Dir.Open(name)
}
