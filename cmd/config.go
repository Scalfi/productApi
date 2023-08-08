package main

import (
	httpGo "net/http"
	"prodocutApi/routes"
)

func Start() error {
	routers := routes.NewHandler()
	return httpGo.ListenAndServe(":2345", routers)
}
