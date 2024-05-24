package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Login(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	ip := getIpAddress(req)

	// TODO: Implement proper authentication with the database instead.
	authenticated := username == "test" && password == "test123"

	if authenticated {
		http.Redirect(res, req, "/admin", http.StatusTemporaryRedirect)
		// TODO: Implant the auth token into the client's cookie.

		if ip != "" {
			log.Printf("User %q logged in from %s\n", username, ip)
		}
	} else {
		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(res, `<div id="error-box">Incorrect credentials.</div>`)

		if ip == "" {
			log.Printf("Failed login attempt: %q and %q\n", username, password)
		} else {
			// TODO: Use the IP to implement anti-spam.
			log.Printf("Failed login attempt: %q and %q from %s\n", username, password, ip)
		}
	}
}

func getIpAddress(req *http.Request) string {
	const IP_SEPARATOR = ", "

	ip := (func() string { // Pseudo match statement.
		if ip := req.Header.Get("X-Real-Ip"); ip != "" {
			return ip
		}
		if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
			return ip
		}
		if ip := req.RemoteAddr; ip != "" {
			return ip
		}
		return ""
	})()

	ips := strings.Split(ip, IP_SEPARATOR)
	return ips[0]
}
