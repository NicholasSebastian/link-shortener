package main

import (
	// "database/sql"
	"fmt"
	"link-shortener/auth"
	"net/http"
)

// TODO: Database implementation.
// TODO: Page with form and HTMX implementation.

func main() {
	host := "localhost:8080"
	linkShortener := NewLinkShortener(host)

	directory := http.Dir("./static")
	fileServer := http.FileServer(directory)

	authRouter := http.NewServeMux()
	authRouter.Handle("GET /admin", fileServer)
	authRouter.HandleFunc("POST /shorten", linkShortener.Shorten)

	router := http.NewServeMux()
	router.Handle("/", auth.Middleware(authRouter))
	router.Handle("GET /", fileServer)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /{path}", linkShortener.Redirect)
	router.HandleFunc("POST /login", auth.Login)

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
