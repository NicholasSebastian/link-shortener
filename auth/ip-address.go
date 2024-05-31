package auth

import (
	"link-shortener/utils"
	"net/http"
	"strings"
)

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
