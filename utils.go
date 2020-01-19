package main

import (
	"net/http"
	"strings"
)

// Error - represents an error
type Error struct {
	message string
}

func (error *Error) Error() string {
	return error.message
}

// GetClientIPAddress - Gets the client IP address from the request headers
func GetClientIPAddress(request *http.Request) string {

	if request == nil {
		return ""
	}

	forwarded := request.Header.Get("X-FORWARDED-FOR")

	if len(forwarded) == 0 {
		return request.RemoteAddr
	}

	// forwarded may contain multiple IPs, get the first one
	ips := strings.Split(forwarded, ", ")

	return ips[0]
}
