package handlers

import "net/http"

func Writer(w http.ResponseWriter, r *http.Request) {}

func Reader(w http.ResponseWriter, r *http.Request) {}

func IniterHandlers() {
	http.HandleFunc("/", Reader)
	http.HandleFunc("/write", Writer)
}
