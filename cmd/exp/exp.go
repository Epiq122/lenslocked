package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.tmpl")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Rob Gleason",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
