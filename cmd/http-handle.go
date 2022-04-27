package main

import (
	"log"
	"net/http"
)

type StatusRespWr struct {
	http.ResponseWriter
	status int
}

func (w *StatusRespWr) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

// wrapHandler wraps a http.HandlerFunc up so it logs any errors, such as 404.
func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srw := &StatusRespWr{ResponseWriter: w}
		log.Printf("Serving %s", r.RequestURI)
		h.ServeHTTP(srw, r)
		// 400+ codes are the error codes, so only log if there was an error
		if srw.status >= 400 {
			log.Printf("Error status code: %d when serving path: %s",
				srw.status, r.RequestURI)
		}
	}
}

// enableCors is used to enable CORS on a http.HandlerFunc.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
