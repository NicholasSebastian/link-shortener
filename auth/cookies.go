package auth

import "net/http"

func setAuthCookie(res *http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:  "",
		Value: token,
		// TODO
	}

	http.SetCookie(*res, &cookie)
}

func removeAuthCookie(res *http.ResponseWriter) {
	// TODO
}
