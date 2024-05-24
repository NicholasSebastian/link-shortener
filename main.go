package main

import (
	// "database/sql"
	"fmt"
	"link-shortener/auth"
	"link-shortener/shortener"
	"log"
	"net/http"
)

// TODO: Database implementation.
// TODO: Page with form and HTMX implementation.

func main() {
	// database, err := sql.Open("sqlite3", "./links.db")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer database.Close()

	directory := htmlDir{
		dir: http.Dir("./static"),
	}

	router := http.NewServeMux()
	fileServer := http.FileServer(directory)
	linkShortener := shortener.New()

	// Regular routes.
	router.Handle("GET /", fileServer)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /{path}", linkShortener.Redirect)
	router.HandleFunc("POST /login", auth.Login)

	// Authenticated routes.
	router.Handle("GET /admin", auth.Middleware(fileServer))
	router.Handle("POST /admin", auth.Middleware(fileServer)) // To handle redirects from logins.
	router.Handle("POST /shorten", auth.MiddlewareFunc(linkShortener.Shorten))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is running on %s\n", server.Addr)
	err := server.ListenAndServe()
	log.Fatalln(err)
}
