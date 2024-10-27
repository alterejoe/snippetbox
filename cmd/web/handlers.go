package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	fmt.Fprint(w, "Root")
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	viewid := r.PathValue("id")
	id, err := strconv.Atoi(viewid)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display specific snippet with the ID: %d ...", id)
	// msg := fmt.Sprintf("Display specific snipped with the ID: %d ...", id)
	// w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form to create a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Create a new snippet")
}
