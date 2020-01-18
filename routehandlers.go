package main

import (
	"html"
	"fmt"
	"regexp"
	"net/http"
)

type route struct {
	urlPattern *regexp.Regexp
	handler http.Handler
}

type githubHandler struct { }
func (handler githubHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from GitHub! %q", html.EscapeString(request.URL.Path))
}

type gitlabHandler struct { }
func (handler gitlabHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from GitLab! %q", html.EscapeString(request.URL.Path))
}

type bitbucketHandler struct { }
func (handler bitbucketHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hello from BitBucket! %q", html.EscapeString(request.URL.Path))
}

var GitHubUrlPattern = `^/github/([^/]+)/([^/]+)/?$`
var GitLabUrlPattern = `^/gitlab/([^/]+)/([^/]+)/?$`
var BitBucketUrlPattern = `^/bitbucket/([^/]+)/([^/]+)/?$`

func GetGitHubHandler() *route {
	var githubRegex *regexp.Regexp = regexp.MustCompile(GitHubUrlPattern)
	return &route{githubRegex, githubHandler{}}
}

func GetGitLabHandler() *route {
	var gitlabRegex *regexp.Regexp = regexp.MustCompile(GitLabUrlPattern)
	return &route{gitlabRegex, gitlabHandler{}}
}

func GetBitBucketHandler() *route {
	var bitbucketRegex *regexp.Regexp = regexp.MustCompile(BitBucketUrlPattern)
	return &route{bitbucketRegex, bitbucketHandler{}}
}
