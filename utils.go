package main

// Error - represents an error
type Error struct {
	message string
}

func (error *Error) Error() string {
	return error.message
}
