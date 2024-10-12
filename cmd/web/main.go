package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// file server which serves files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
