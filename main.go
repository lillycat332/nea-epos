package main

import (
	"fmt"
)

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