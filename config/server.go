package config

import (
	"api-std-packages/pkg/middleware"
	"api-std-packages/pkg/routes"
	"net/http"
)

var router routes.Routes

func ConfigServer() {
	router = append(router, routes.AppRoutes()...)
	router = append(router, routes.ProductRoutes()...)

	for _, route := range router {
		var handler http.Handler

		if route.Auth {
			handler = middleware.AuthMiddleware(route.Handler)
		} else {
			handler = route.Handler
		}

		if route.Group != "" {
			http.Handle(route.Group + route.Path, handler)
			continue
		}
		
		http.Handle(route.Path, handler)
	}
}