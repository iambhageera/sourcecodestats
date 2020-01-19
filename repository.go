package main

import (
	"log"
	"strings"
)

// Repository - Stores the information about the repository
type Repository struct {
	name    string
	service string
	owner   string
}

// ParseURL - parses the request URL to get the repository information
func (repo *Repository) ParseURL(url string) (bool, *Error) {

	if len(url) == 0 {
		return false, &Error{"URL cannot be empty!"}
	}

	// Remove the forward slash from the beginning to avoid edge cases
	if url[0] == byte('/') {
		url = url[1:]
	}

	var tokens []string = strings.Split(url, "/")

	switch len(tokens) {
	case 1:
		return false, &Error{"Username and Repository is missing!"}
	case 2:
		return false, &Error{"Repository name is missing!"}
	default:
		log.Printf("Received valid request path - [%s]", url)
	}

	// Store the repository details
	repo.service = tokens[0]
	repo.owner = tokens[1]
	repo.name = tokens[2]

	return true, nil
}
