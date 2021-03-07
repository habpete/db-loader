package main

import (
	"db-loader/handlers"
	"db-loader/writer"
	"net/http"
)

func main() {
	dbw := &writer.DBWriter{}
	handlers.IniterHandlers(dbw)
	http.ListenAndServe(":8080", nil)
}
