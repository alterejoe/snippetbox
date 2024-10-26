package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// 38570
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Root")
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	viewid := r.PathValue("id")
	id, err := strconv.Atoi(viewid)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display specific snipped with the ID: %d ...", id)
	// msg := fmt.Sprintf("Display specific snipped with the ID: %d ...", id)
	// w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet create"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(http.StatusCreated)
	w.Header().Add("Server", "Go")
	io.WriteString(w, "Snippet create post is working")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", root)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create/", snippetCreate)
	mux.HandleFunc("POST /snippet/create/", snippetCreatePost)

	log.Print("Starting on :3000")

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
	log.Print("Server stopped")
}
