package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the danger zones!</h1>")
}

func contactHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">robgleason122@gmail.com</a>.</p>")
}

type Router struct{}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		handlerFunc(w, req)
	case "/contact":
		contactHandler(w, req)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	r := Router{}
	fmt.Println("Starting server on port :3000...")
	http.ListenAndServe(":3000", r)
}
