package auth

import (
	"fmt"
	"link-shortener/utils"
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
		res.Header().Set("HX-Redirect", "/admin")

		// TODO: Implant the auth token into the client's cookie.

		if ip != "" {
			log.Printf("User %q logged in from %s\n", username, ip)
		}
	} else {
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
	ip := utils.Fallback(
		req.Header.Get("X-Real-Ip"),
		req.Header.Get("X-Forwarded-For"),
		req.RemoteAddr,
	)

	const IP_SEPARATOR = ", "
	ips := strings.Split(ip, IP_SEPARATOR)

	return ips[0]
}
