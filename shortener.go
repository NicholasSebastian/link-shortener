package main

import (
	"fmt"
	"net/http"
)

type LinkShortener struct {
	host string
	urls map[string]string
}

func NewLinkShortener(host string) *LinkShortener {
	return &LinkShortener{
		host: host,
		urls: make(map[string]string),
	}
}

func (ls *LinkShortener) Shorten(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	original := req.FormValue("url")
	key := req.FormValue("path")

	if len(original) == 0 || len(key) == 0 {
		http.Error(res, "Missing form value(s).", http.StatusBadRequest)
		return
	}

	// TODO: Add this to a database instead.
	ls.urls[key] = original

	newUrl := fmt.Sprintf("%s/%s", ls.host, key)
	responseHtml := fmt.Sprintf(`
		<div>Original URL: %s</div>
		<div>Shortened URL: %s</div>
	`, original, newUrl)

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, responseHtml)
}

func (ls *LinkShortener) Redirect(res http.ResponseWriter, req *http.Request) {
	key := req.PathValue("path")
	if len(key) == 0 {
		http.Error(res, "Invalid path.", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from the database instead.
	target, exists := ls.urls[key]

	if !exists {
		http.Error(res, "Path does not exist.", http.StatusNotFound)
		return
	}

	// TODO: This is redirecting incorrectly. For some reason it redirects to a path.
	http.Redirect(res, req, target, http.StatusMovedPermanently)
}
