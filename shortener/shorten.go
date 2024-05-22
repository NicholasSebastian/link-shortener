package shortener

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var pathRegex, _ = regexp.Compile(`^/?[\w-]+$`)

func (ls *LinkShortener) Shorten(res http.ResponseWriter, req *http.Request) {
	original := req.FormValue("url")
	key := req.FormValue("path")

	if len(original) == 0 || len(key) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "<div id='error-box'>Missing form value(s).</div>")
		return
	}
	if _, err := url.ParseRequestURI(original); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "<div id='error-box'>Invalid URL.</div>")
		return
	}
	if !pathRegex.MatchString(key) {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "<div id='error-box'>Invalid Path.</div>")
		return
	}

	// TODO: Add this to a database instead.
	ls.urls[key] = original
	log.Printf("New path created at /%s", key)

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, "<div>Original URL: %s</div><div>New Path: /%s</div>", original, key)
}
