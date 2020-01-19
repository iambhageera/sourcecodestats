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

	// Verified that request is valid.
	// Processing the request may take few seconds to many minutes based
	// on the size of the source code. Letting the client wait without any
	// notification is a terrible user experience. Walk the client through
	// the steps while processing the source code in the repository.
	// Server Side Events can be used to acheive this.

	// Initiate the SSE Registrar which can send regular updates to the client
	var registrar = ServerSideEventRegistrar{
		messageCounter: 0,
		request:        request,
		connection:     responseWriter,
	}

	// Prepare the connection to handle the events
	if ok, err := registrar.PrepareConnection(); !ok {
		http.Error(responseWriter, err.message, http.StatusInternalServerError)
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
