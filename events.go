package main

import (
	"fmt"
	"net/http"
)

// ServerSideEventRegistrar - registers the server side events with the client
type ServerSideEventRegistrar struct {
	messageCounter int
	request        *http.Request
	responseWriter http.ResponseWriter
}

// RegisterEvent - Method to register or emit a new server side event to the client
func (registrar *ServerSideEventRegistrar) RegisterEvent(eventType string, message string) {
	registrar.messageCounter++
	fmt.Print(registrar.responseWriter, "id: %d\n", registrar.messageCounter)
	fmt.Print(registrar.responseWriter, "event: %s\n", eventType)
	fmt.Print(registrar.responseWriter, "data: %s\n\n", message)
}
