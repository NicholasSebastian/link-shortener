package auth

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		log.Printf("Authorization: %q", authHeader) // For debugging only.

		// TODO: Implementation.
		next.ServeHTTP(res, req)
	})
}

func MiddlewareFunc(next http.HandlerFunc) http.Handler {
	nextFunc := http.HandlerFunc(next)
	return Middleware(nextFunc)
}
