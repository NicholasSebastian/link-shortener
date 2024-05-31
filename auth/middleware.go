package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func parseToken(tokenstr string) (*jwt.Token, error) {
	return jwt.Parse(tokenstr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			signMethod := t.Header["alg"]
			err := fmt.Errorf("unexpected signing method: %v", signMethod)
			return nil, err
		}
		return []byte(secretKey), nil
	})
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		tokenstr := req.Header.Get("Authorization")
		log.Printf("Authorization: %q", tokenstr) // For debugging only.

		token, err := parseToken(tokenstr)
		if err != nil {
			fmt.Fprint(res, err.Error())
			res.WriteHeader(http.StatusBadRequest)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username := claims["username"]
			log.Println(username) // For debugging only.
		} else {
			fmt.Fprint(res, err.Error())
			res.WriteHeader(http.StatusBadRequest)
		}

		// TODO: Implement authorization checking.

		next.ServeHTTP(res, req)
	})
}

func MiddlewareFunc(next http.HandlerFunc) http.Handler {
	nextFunc := http.HandlerFunc(next)
	return Middleware(nextFunc)
}
