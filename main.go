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
	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home-page.tmpl", "tailwind.tmpl"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact-page.tmpl", "tailwind.tmpl"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq-page.tmpl", "tailwind.tmpl"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
