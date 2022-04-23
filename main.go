package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// This file is just for initializing the program - everything else is in the other files

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `servepls is a simple webserver.
		Launching this will activate the webserver on port :8080 or a port specified by -port.
		For more information see the documentation.
		Usage:
		`)
		flag.PrintDefaults()
	}
}

func main() {
	port := flag.String("port", ":8080", "a port number prefixed by :, that tells the program where it should host the server.\nin most situations, you do not need to change this.")
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
