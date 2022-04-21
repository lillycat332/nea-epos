package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "modernc.org/sqlite"
)

type StatusRespWr struct {
	http.ResponseWriter
	status int
}

func (w *StatusRespWr) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func main() {
	port := flag.String("port", ":8080", "HTTP Port Number")
	flag.Parse()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", wrapHandler(fileServer))
	http.HandleFunc("/form", formHandler)
	log.Printf("Starting server on port %s\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	log.Printf("POST request %s, %s recieved \n", r.RemoteAddr, r.Form)
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintf(w, "username = %s\n", username)
	fmt.Fprintf(w, "password = %s\n", password)
}

func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srw := &StatusRespWr{ResponseWriter: w}
		log.Printf("Serving %s", r.RequestURI)
		h.ServeHTTP(srw, r)
		if srw.status >= 400 { // 400+ codes are the error codes
			log.Printf("Error status code: %d when serving path: %s",
				srw.status, r.RequestURI)
		}
	}
}
