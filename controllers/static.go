package controllers

import (
	"html/template"
	"net/http"

	"github.com/epiq122/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Will there be a free version?",
			Answer:   "Yes! we offer a free 30 day free trial",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have a support staff answering emails 24hrs a day",
		},
		{
			Question: "What's your boogle?",
			Answer:   "No boggle here!",
		},
		{
			Question: "Do you have an office?",
			Answer:   "Nope!!",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
