package main

// Error - represents an error
type Error struct {
	message string
}

func (error *Error) Error() string {
	return error.message
}

// Repository - Stores the information about the
type Repository struct {
	name    string
	service string
	owner   string
}
