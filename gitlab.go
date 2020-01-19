package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
)

// GitLabURLPattern - Regex Pattern for GitLab request URLs
var GitLabURLPattern = `^/gitlab/([^/]+)/([^/]+)/?$`

// GitlabHandler - Request handler type for GitLab
type GitlabHandler struct{}

func (handler GitlabHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	var registrar = ServerSideEventRegistrar{
		messageCounter: 0,
		request:        request,
		connection:     responseWriter,
	}

	if ok, err := registrar.PrepareConnection(); !ok {
		http.Error(responseWriter, err.message, http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(responseWriter, "Hello from GitLab! %q", html.EscapeString(request.URL.Path))
}

// GetGitLabHandler - Gets the route handler type for GitLab
func GetGitLabHandler() *Route {
	var gitlabRegex *regexp.Regexp = regexp.MustCompile(GitLabURLPattern)
	return &Route{gitlabRegex, GitlabHandler{}}
}
