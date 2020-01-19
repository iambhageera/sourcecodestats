package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
)

// GitHubURLPattern - Regex Pattern for GitHub request URLs
var GitHubURLPattern = `^/github/([^/]+)/([^/]+)/?$`

// GithubHandler - Request handler type for GitHub
type GithubHandler struct{}

func (handler GithubHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	var registrar = ServerSideEventRegistrar{
		messageCounter: 0,
		request:        request,
		connection:     responseWriter,
	}

	if ok, err := registrar.PrepareConnection(); !ok {
		http.Error(responseWriter, err.message, http.StatusInternalServerError)
		return
	}

	var repo Repository = Repository{}

	if ok, _ := repo.ParseURL(request.URL.Path); !ok {
		http.NotFound(responseWriter, request)
		return
	}

	fmt.Fprintf(responseWriter, "Hello from GitHub! %q", html.EscapeString(request.URL.Path))
}

// GetGitHubHandler - Gets the route handler for GitHub
func GetGitHubHandler() *Route {
	var githubRegex *regexp.Regexp = regexp.MustCompile(GitHubURLPattern)
	return &Route{githubRegex, GithubHandler{}}
}
