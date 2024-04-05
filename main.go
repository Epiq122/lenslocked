package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/epiq122/lenslocked/controllers"
	"github.com/epiq122/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func exectuteTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was a error parsing the template", http.StatusInternalServerError)

	}
	tpl.Execute(w, nil)
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
	homeTpl, err := views.Parse(filepath.Join("templates", "home.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Method(http.MethodGet, "/", controllers.Static{
		Template: homeTpl,
	})
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
