package main

import (
	// "database/sql"
	"fmt"
	"link-shortener/auth"
	"link-shortener/shortener"
	"net/http"
)

// TODO: Database implementation.
// TODO: Page with form and HTMX implementation.

func main() {
	host := "localhost:8080"
	router := http.NewServeMux()

	directory := http.Dir("./static")
	fileServer := http.FileServer(directory)
	linkShortener := shortener.New(host)

	// Regular routes.
	router.Handle("GET /", fileServer)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /{path}", linkShortener.Redirect)
	router.HandleFunc("GET /login", auth.Login) // This is a 'GET' request because it will redirect to "/admin".

	// Authenticated routes.
	router.Handle("POST /shorten", auth.Middleware(linkShortener.Shorten))
	router.Handle("GET /admin", auth.Middleware(func(w http.ResponseWriter, r *http.Request) {
		// We handle this route manually because we want to strip away the ".html" from the path.
		r.URL.Path = "admin.html"
		fileServer.ServeHTTP(w, r)
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is running on %s\n", host)
	server.ListenAndServe()

	// database, err := sql.Open("sqlite3", "./links.db")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer database.Close()
}
