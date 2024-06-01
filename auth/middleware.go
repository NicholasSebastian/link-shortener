package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		signMethod := token.Header["alg"]
		err := fmt.Errorf("unexpected signing method: %v", signMethod)
		return nil, err
	}
	return []byte(secretKey), nil
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenstr := req.Header.Get("Authorization")
		log.Printf("Authorization: %q", tokenstr) // For debugging only.

		token, err := jwt.Parse(tokenstr, keyFunc)
		if err != nil {
			fmt.Fprint(res, err.Error())
			res.WriteHeader(http.StatusBadRequest)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Fprint(res, "Failed to parse session data.")
			res.WriteHeader(http.StatusInternalServerError)
			// TODO: Redirect back to login page.
		}

		username := claims["username"]
		log.Println(username) // For debugging only.
		// TODO: Implement authorization checking.

		next.ServeHTTP(res, req)
	})
}

func MiddlewareFunc(next http.HandlerFunc) http.Handler {
	nextFunc := http.HandlerFunc(next)
	return Middleware(nextFunc)
}
