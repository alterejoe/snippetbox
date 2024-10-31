package main

import (
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.RequestURI))
		// log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error(), slog.String("method", r.Method), slog.String("uri", r.RequestURI))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
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

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display form to create a new snippet"))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Create a new snippet")
}
