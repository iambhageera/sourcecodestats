package main

import (
	"fmt"
	"net/http"
)

// ServerSideEventRegistrar - registers the server side events with the client
type ServerSideEventRegistrar struct {
	messageCounter int
	request        *http.Request
	connection     http.ResponseWriter
}

// PrepareConnection - Sets up the client connection to handle event streaming
func (registrar *ServerSideEventRegistrar) PrepareConnection() (bool, *Error) {

	if registrar.connection != nil {

		// Make sure that the connection supports flushing
		if _, ok := registrar.connection.(http.Flusher); !ok {
			return false, &Error{"Streaming Unsupported!"}
		}

		// Set all the required HTTP headers for handling Server Side Events
		registrar.connection.Header().Set("Content-Type", "text/event-stream")
		registrar.connection.Header().Set("Cache-Control", "no-cache")
		registrar.connection.Header().Set("Connection", "keep-alive")
		registrar.connection.Header().Set("Access-Control-Allow-Origin", "*")
	}

	return true, nil
}

// RegisterEvent - Method to register or emit a new server side event to the client
func (registrar *ServerSideEventRegistrar) RegisterEvent(eventType string, message string) {
	registrar.messageCounter++
	fmt.Fprintf(registrar.connection, "id: %d\n", registrar.messageCounter)
	fmt.Fprintf(registrar.connection, "event: %s\n", eventType)
	fmt.Fprintf(registrar.connection, "data: %s\n\n", message)
}
