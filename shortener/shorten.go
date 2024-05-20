package shortener

import (
	"fmt"
	"log"
	"net/http"
)

func (ls *LinkShortener) Shorten(res http.ResponseWriter, req *http.Request) {
	original := req.FormValue("url")
	key := req.FormValue("path")

	if len(original) == 0 || len(key) == 0 {
		http.Error(res, "Missing form value(s).", http.StatusBadRequest)
		return
	}

	// TODO: Add this to a database instead.
	ls.urls[key] = original
	log.Fatalf("New path created at %s/%s", ls.host, key)

	newUrl := fmt.Sprintf("%s/%s", ls.host, key)
	responseHtml := fmt.Sprintf(`
		<div>Original URL: %s</div>
		<div>Shortened URL: %s</div>
	`, original, newUrl)

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprint(res, responseHtml)
}
