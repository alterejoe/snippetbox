package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := path + "/index.html"
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}

func main() {
	mux := http.NewServeMux()
	addr := flag.String(("addr"), ":3000", "HTTP network address")
	_ = flag.String("unique-id", "", "Unique ID for development")
	flag.Parse()

	fileserver := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	log.Print("Server started")

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create/", snippetCreate)
	mux.HandleFunc("POST /snippet/create/", snippetCreatePost)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err)
	log.Print("Server stopped")
}
