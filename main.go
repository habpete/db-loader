package main

import (
	"db-loader/handlers"
	"net/http"
)

func main() {
	handlers.IniterHandlers()
	http.ListenAndServe(":8080", nil)
}
