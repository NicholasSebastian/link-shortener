package auth

import (
	"log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		log.Println(authHeader) // For debugging only.

		// TODO
		next.ServeHTTP(res, req)
	})
}
