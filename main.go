package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":8080", "Service Port Number")
	flag.Parse()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", wrapHandler(fileServer))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", userCreateHandler)
	log.Printf("Starting server on port %s\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal(err)
	}
}
