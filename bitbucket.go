package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
)

// BitBucketURLPattern - Regex Pattern for BitBucket request URLs
var BitBucketURLPattern = `^/bitbucket/([^/]+)/([^/]+)/?$`

// BitbucketHandler - Request handler for BitBucket
type BitbucketHandler struct{}

func (handler BitbucketHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	var registrar = ServerSideEventRegistrar{
		messageCounter: 0,
		request:        request,
		connection:     responseWriter,
	}

	if ok, err := registrar.PrepareConnection(); !ok {
		http.Error(responseWriter, err.message, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(responseWriter, "Hello from BitBucket! %q", html.EscapeString(request.URL.Path))
}

// GetBitBucketHandler - Gets the route handler for BitBucket
func GetBitBucketHandler() *Route {
	var bitbucketRegex *regexp.Regexp = regexp.MustCompile(BitBucketURLPattern)
	return &Route{bitbucketRegex, BitbucketHandler{}}
}
