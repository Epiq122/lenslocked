package controllers

import (
	"net/http"
)

type Template interface {
	Execute(w http.ResponseWriter, data interface{})
}

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)

}
