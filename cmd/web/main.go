package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// input arg
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// log settings
	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// file server which serves files out of "./ui/static"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	// use fatal and panic only in main function
	errorLog.Fatal(err)
}
