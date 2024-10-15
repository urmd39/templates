package helpers

import (
	"net"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func GenCode() string {
	tempId := uuid.New().String()
	code := tempId[0:8]
	return strings.ToUpper(code)
}

func GetRealAddr(r *http.Request) (remoteIP, message string) {
	remoteIP = ""
	message = ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
		message += ";" + parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		message += ";" + lastFwd
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
			message += ";" + ip.String()
		}
	}
	return remoteIP, message
}
