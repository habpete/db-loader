package handlers

import (
	"db-loader/writer"
	"net/http"
)

type HandlersState struct {
	wr writer.IWriter
}

func (this *HandlersState) Writer(w http.ResponseWriter, r *http.Request) {
	this.wr.Write(nil)
}

func (this *HandlersState) Reader(w http.ResponseWriter, r *http.Request) {}

func IniterHandlers(wr writer.IWriter) {
	hs := &HandlersState{
		wr: wr,
	}
	http.HandleFunc("/", hs.Reader)
	http.HandleFunc("/write", hs.Writer)
}
