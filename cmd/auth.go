package main

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	log.Printf("POST request (Login) recieved from %s", r.RemoteAddr)
	log.Fatal("I haven't implemented this yet!")
}
