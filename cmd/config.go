package main

import (
	httpGo "net/http"
	"prodocutApi/http"
)

func Start() error {
	routers := http.NewHandler()
	return httpGo.ListenAndServe(":2345", routers)
}
