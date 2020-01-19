package main

import (
	"fmt"
	"html"
	"net/http"
	"regexp"
)

// GitHubURLPattern - Regex Pattern for GitHub request URLs
var GitHubURLPattern = `^/github/([^/]+)/([^/]+)/?$`

// GithubRequestHandler - Request handler type for GitHub
type GithubRequestHandler struct{}

func (handler GithubRequestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	var registrar = ServerSideEventRegistrar{
		messageCounter: 0,
		request:        request,
		connection:     responseWriter,
	}

	if ok, err := registrar.PrepareConnection(); !ok {
		http.Error(responseWriter, err.message, http.StatusInternalServerError)
		return
	}

	var repo GitHubRepository = GitHubRepository{}

	if ok, _ := repo.ParseRequestURL(request.URL.Path); !ok {
		http.NotFound(responseWriter, request)
		return
	}

	// Generate the repo URLs
	repo.MakeURL()

	if ok, err := repo.VerifyRepository(); !ok {
		http.Error(responseWriter, err.message, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(responseWriter, "Hello from GitHub! %q", html.EscapeString(request.URL.Path))
}

// GitHubRepository -
type GitHubRepository struct {
	Repository
}

// MakeURL - Generates the URL for the repository
func (repo *GitHubRepository) MakeURL() (bool, *Error) {

	if len(repo.service) == 0 || len(repo.owner) == 0 || len(repo.name) == 0 {
		return false, &Error{"Required repository information is missing!"}
	}

	repo.ownerURL = "https://github.com/" + repo.owner
	repo.url = repo.ownerURL + "/" + repo.name

	return true, nil
}

// GetGitHubHandler - Gets the route handler for GitHub
func GetGitHubHandler() *Route {
	var githubRegex *regexp.Regexp = regexp.MustCompile(GitHubURLPattern)
	return &Route{githubRegex, GithubRequestHandler{}}
}
