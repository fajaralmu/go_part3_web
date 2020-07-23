package main

import (
	"net/http"
	"text/template"
)

//App is the web app
type App struct {
}

func (app *App) start() {
	println("__will start app__")
	http.HandleFunc("/", app.homeRoute)
	http.ListenAndServe(":8080", nil)
}

type homePageData struct {
	Title   string
	Message string
}

func (app *App) homeRoute(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")

	if err == nil {
		pageData := homePageData{
			Title:   "Welcome",
			Message: "Hello World",
		}
		tmpl.Execute(w, pageData)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))

	}
}

func main() {
	var app App = App{}
	app.start()
}
