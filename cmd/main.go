package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// This file is just for initializing the program - everything else is in the other files

var (
	db   string
	port string
	fs   string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage:`)
		flag.PrintDefaults()
	}
	port = *flag.String("port", ":8080", "a port number prefixed by :, that tells the program where it should host the server.")
	db = *flag.String("db", "./EPOS.db", "path to an sqlite3 database file, that tells the program where to store and load it's data from.")
	fs = *flag.String("fs", "./static", "path to your html files.")
	flag.Parse()
}

func main() {
	fileServer := http.FileServer(http.Dir(fs))
	http.Handle("/", wrapHandler(fileServer))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/create", userCreateHandler)
	http.HandleFunc("/readUsers", userReader)
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
