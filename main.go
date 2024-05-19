package main

import (
	// "database/sql"
	"fmt"
	"net/http"
)

// TODO: Database implementation.
// TODO: Page with form and HTMX implementation.

func main() {
	host := "localhost:8080"
	linkShortener := NewLinkShortener(host)

	directory := http.Dir("./static")
	fileServer := http.FileServer(directory)

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/shorten", linkShortener.Shorten)
	http.HandleFunc("/{path}", linkShortener.Redirect)
	http.Handle("/", fileServer)

	fmt.Printf("Server is running on %s\n", host)
	http.ListenAndServe(":8080", nil)

	// database, err := sql.Open("sqlite3", "./links.db")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer database.Close()
}
