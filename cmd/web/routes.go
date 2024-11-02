package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// create server mux
	mux := http.NewServeMux()

	// file server which serves files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// register handlers
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
