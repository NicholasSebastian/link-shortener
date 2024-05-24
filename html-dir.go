package main

import (
	"net/http"
	"path/filepath"
)

// Overrides the default filesystem behavior by falling back to '.html'.
// I don't know why this isn't the default behaviour for a HTTP static file server.

type htmlDir struct {
	dir http.Dir
}

func (fs htmlDir) Open(name string) (http.File, error) {
	if name != "/" && filepath.Ext(name) == "" {
		name += ".html"
	}
	return fs.dir.Open(name)
}
