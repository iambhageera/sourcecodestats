package main

import (
	"net/http"
	"regexp"
)

// Route - Structure to store details of a route
type Route struct {
	urlPattern *regexp.Regexp
	handler    http.Handler
}

// HTTPRouteHandleManager - Manages all the routes and their handlers
type HTTPRouteHandleManager struct {
	routes []*Route
}

// UniversalHandler - A single request handler which will process all the requests for the application
func (handleManager *HTTPRouteHandleManager) UniversalHandler(responseWriter http.ResponseWriter, request *http.Request) {

	// Stores the final handler for the request
	var handler http.Handler

	for _, route := range handleManager.routes {
		if ok := route.urlPattern != nil && route.handler != nil; ok && route.urlPattern.MatchString(request.URL.Path) {
			handler = route.handler
			break
		}
	}

	if handler == nil {
		http.NotFound(responseWriter, request)
	}

	handler.ServeHTTP(responseWriter, request)
}

// InitializeRouteHandles - Initializes the handles for all available routes
func (handleManager *HTTPRouteHandleManager) InitializeRouteHandles() func(responseWriter http.ResponseWriter, request *http.Request) {

	handleManager.routes = append(handleManager.routes, GetGitHubHandler())
	handleManager.routes = append(handleManager.routes, GetGitLabHandler())
	handleManager.routes = append(handleManager.routes, GetBitBucketHandler())

	return handleManager.UniversalHandler
}

// InitializeRouteHandleManager - Initializes the HttpRouteHandleManager object
func InitializeRouteHandleManager() *HTTPRouteHandleManager {
	var routeHandler *HTTPRouteHandleManager = new(HTTPRouteHandleManager)
	routeHandler.routes = make([]*Route, 0)
	return routeHandler
}
