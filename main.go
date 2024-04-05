package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func exectuteTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was a error parsing the template", http.StatusInternalServerError)

	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was a error parsing the template", http.StatusInternalServerError)

	}
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	tplPath := filepath.Join("templates", "home.tmpl")
	exectuteTemplate(w, tplPath)

}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	tplPath := filepath.Join("templates", "contact.tmpl")
	exectuteTemplate(w, tplPath)

}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	tplPath := filepath.Join("templates", "faq.tmpl")
	exectuteTemplate(w, tplPath)

}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
