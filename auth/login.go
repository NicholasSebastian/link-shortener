package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "" // TODO: Retrieve the secret key from the environment variables instead.

func Login(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	ip := getIpAddress(req)

	// TODO: Implement proper authentication with the database instead.
	authenticated := username == "test" && password == "test123"

	if authenticated {
		claims := jwt.MapClaims{
			"username": username,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenstr, err := token.SignedString(secretKey)
		if err != nil {
			// TODO: Respond with a server error.
			return
		}

		setAuthCookie(&res, tokenstr)
		res.Header().Set("HX-Redirect", "/admin")

		if ip != "" {
			log.Printf("User %q logged in from %s\n", username, ip)
		}
	} else {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(res, `<div>Incorrect credentials.</div>`)

		if ip == "" {
			log.Printf("Failed login attempt: %q and %q\n", username, password)
		} else {
			// TODO: Use the IP to implement anti-spam.
			log.Printf("Failed login attempt: %q and %q from %s\n", username, password, ip)
		}
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	removeAuthCookie(&res)
	// TODO: Redirect the client to the login page.
}
