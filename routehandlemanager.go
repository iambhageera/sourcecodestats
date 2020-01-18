package main

import (
	"net/http"
)

type HttpRouteHandleManager struct {
	routes []*route
}

func (handleManager *HttpRouteHandleManager) UniversalHandler(responseWriter http.ResponseWriter, request *http.Request) {
	
	for _, route := range handleManager.routes {
		if ok := route.urlPattern != nil && route.handler != nil;
			ok && route.urlPattern.MatchString(request.URL.Path) {

			route.handler.ServeHTTP(responseWriter, request)
			return
		}
	}

	http.NotFound(responseWriter, request)
}

func (handleManager *HttpRouteHandleManager) InitializeRouteHandles() func(responseWriter http.ResponseWriter, request *http.Request) {

	handleManager.routes = append(handleManager.routes, GetGitHubHandler())
	handleManager.routes = append(handleManager.routes, GetGitLabHandler())	
	handleManager.routes = append(handleManager.routes, GetBitBucketHandler())

	return handleManager.UniversalHandler
}

func InitializeRouteHandleManager() *HttpRouteHandleManager {
	var routeHandler *HttpRouteHandleManager = new(HttpRouteHandleManager)
	routeHandler.routes = make([]*route, 0)
	return routeHandler
}
