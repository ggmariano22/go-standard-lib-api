package routes

import (
	"api-std-packages/pkg/handlers"
	"net/http"
)

type Route struct {
	Name string
	Group string
	Path string
	Auth bool
	Handler http.Handler
}

type Routes []Route

func AppRoutes() Routes {
	return Routes{
		Route{
			Name: "Health check",
			Group: "",
			Path: "/",
			Auth: false,
			Handler: http.HandlerFunc(handlers.NewHealthCheckHandler().Handle),
		},
	}
}

func ProductRoutes() Routes {
	return Routes{
		Route{
			Name: "Create product",
			Group: "/products",
			Path: "",
			Auth: true,
			Handler: http.HandlerFunc(handlers.NewCreateProductHandler().Handle),
		},
	}
}