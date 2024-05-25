package auth

import "net/http"

func Logout(res http.ResponseWriter, req *http.Request) {
	// TODO: Clear the client's cookie.
	// TODO: Redirect the client to the login page.
}
