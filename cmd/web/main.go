package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// input arg
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// log settings
	infoLog := log.New(os.Stdout, "INFO\t", log.LUTC|log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	// initialize app
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// build custom server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// starting server
	infoLog.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
