package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// This file is just for initializing the program - everything else is in the other files

var (
	db   string
	port string
	fs   string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `servepls is a simple webserver.
Launching this will run it on port :8080 or a port specified by -port.
For more information see the documentation.
Usage:
`)
		flag.PrintDefaults()
	}
}

func main() {
	port = *flag.String("port", ":8080", "a port number prefixed by :, that tells the program where it should host the server..")
	db = *flag.String("db", "./EPOS.db", "path to an sqlite3 database file, that tells the program where to store and load it's data from.")
	fs = *flag.String("fs", "./static", "path to your html files.")
	flag.Parse()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", wrapHandler(fileServer))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", userCreateHandler)
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
