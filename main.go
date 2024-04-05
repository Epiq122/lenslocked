package main

import (
	"fmt"
	"net/http"

	"github.com/epiq122/lenslocked/controllers"
	"github.com/epiq122/lenslocked/templates"
	"github.com/epiq122/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	tpl := views.Must(views.ParseFS(templates.FS, "home.tmpl"))
	r.Get("/", controllers.StaticHandler(tpl))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.tmpl"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.tmpl"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
