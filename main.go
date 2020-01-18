package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Initiliazing App..")

	routeHandleManager := InitializeRouteHandleManager()
	requestHandler := routeHandleManager.InitializeRouteHandles()
	http.HandleFunc("/", requestHandler)

	log.Println("Starting server at port 8080..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
