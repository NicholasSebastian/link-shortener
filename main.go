package main

import (
	// "database/sql"
	"fmt"
	"link-shortener/auth"
	"link-shortener/shortener"
	"link-shortener/utils"
	"log"
	"net/http"
)

// TODO: HTMX implementation.
// TODO: Database implementation.
// TODO: Implement tests.
// TODO: Write the server logs into the database.

func main() {
	// database, err := sql.Open("sqlite3", "./links.db")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer database.Close()

	directory := http.Dir("./static")
	htmlDirectory := utils.HtmlDir{
		Dir: &directory,
	}

	router := http.NewServeMux()
	fileServer := http.FileServer(htmlDirectory)
	linkShortener := shortener.New() // TODO: Pass a reference to the database in here.

	// Regular routes.
	router.Handle("GET /", fileServer)
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /{path}", linkShortener.Redirect)
	router.HandleFunc("GET /logout", auth.Logout)
	router.HandleFunc("POST /login", auth.Login)

	// Authenticated routes.
	router.Handle("GET /admin", auth.Middleware(fileServer))
	router.Handle("POST /shorten", auth.MiddlewareFunc(linkShortener.Shorten))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Server is running on %s\n", server.Addr)
	err := server.ListenAndServe()
	log.Fatalln(err)
}
