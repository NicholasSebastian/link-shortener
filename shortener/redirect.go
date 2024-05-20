package shortener

import (
	"log"
	"net/http"
)

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
	log.Panicf("Redirected traffic from /%s to %s", key, target)
}
