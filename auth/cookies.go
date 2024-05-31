package auth

import "net/http"

func getAuthCookie() string {
	// TODO
	return ""
}

func setAuthCookie(res *http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:  "",
		Value: "",
		// TODO
	}

	http.SetCookie(*res, &cookie)
}
