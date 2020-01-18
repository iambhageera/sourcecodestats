package main

import (
	"fmt"
	"log"
	"net/http"
)

func requestHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Hi there! The URL path is %s!", request.URL.Path)
}

func main() {
	log.Println("An application to count lines of a GitHub repository!")
	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
