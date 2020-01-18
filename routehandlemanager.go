package main

import (
	"net/http"
)

// HTTPRouteHandleManager - Manages all the routes and their handlers
type HTTPRouteHandleManager struct {
	routes []*Route
}

// UniversalHandler - A single request handler which will process all the requests for the application
func (handleManager *HTTPRouteHandleManager) UniversalHandler(responseWriter http.ResponseWriter, request *http.Request) {

	for _, route := range handleManager.routes {
		if ok := route.urlPattern != nil && route.handler != nil; ok && route.urlPattern.MatchString(request.URL.Path) {

			route.handler.ServeHTTP(responseWriter, request)
			return
		}
	}

	http.NotFound(responseWriter, request)
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
