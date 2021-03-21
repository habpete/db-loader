package main

import (
	"db-loader/internal"
	"net/http"
)

func main() {
	internal.New()
	http.ListenAndServe(":8080", nil)
}
