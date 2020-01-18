package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
)

// Route - Structure to store details of a route such as URL pattern
// and it's handler
type Route struct {
	urlPattern *regexp.Regexp
	handler    http.Handler
}

type githubHandler struct{}

func (handler githubHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from GitHub! %q", html.EscapeString(request.URL.Path))
}

type gitlabHandler struct{}

func (handler gitlabHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from GitLab! %q", html.EscapeString(request.URL.Path))
}

type bitbucketHandler struct{}

func (handler bitbucketHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from BitBucket! %q", html.EscapeString(request.URL.Path))
}

// GitHubURLPattern - Regex Pattern for GitHub request URLs
var GitHubURLPattern = `^/github/([^/]+)/([^/]+)/?$`

// GitLabURLPattern - Regex Pattern for GitLab request URLs
var GitLabURLPattern = `^/gitlab/([^/]+)/([^/]+)/?$`

// BitBucketURLPattern - Regex Pattern for BitBucket request URLs
var BitBucketURLPattern = `^/bitbucket/([^/]+)/([^/]+)/?$`

// GetGitHubHandler - Gets the route handler for GitHub
func GetGitHubHandler() *Route {
	var githubRegex *regexp.Regexp = regexp.MustCompile(GitHubURLPattern)
	return &Route{githubRegex, githubHandler{}}
}

// GetGitLabHandler - Gets the route handler for GitLab
func GetGitLabHandler() *Route {
	var gitlabRegex *regexp.Regexp = regexp.MustCompile(GitLabURLPattern)
	return &Route{gitlabRegex, gitlabHandler{}}
}

// GetBitBucketHandler - Gets the route handler for BitBucket
func GetBitBucketHandler() *Route {
	var bitbucketRegex *regexp.Regexp = regexp.MustCompile(BitBucketURLPattern)
	return &Route{bitbucketRegex, bitbucketHandler{}}
}
