package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

//App is the web app
type App struct {
	router *mux.Router
}

func (app *App) start() {

	println("__will start app__")
	http.HandleFunc("/home", app.homeRoute)
	http.ListenAndServe(":8080", nil)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	println("get books")
	writeJSONResponse(w, books)
}

func getIDParams(r *http.Request) string {
	params := mux.Vars(r)
	id := params["id"]
	return id
}

func getBook(w http.ResponseWriter, r *http.Request) {

	writeResponseHeaders(w)
	id := getIDParams(r)
	book := getBookByID(id)

	if nil == book {
		writeErrorMsg(w, "Not Found")
	} else {
		writeJSONResponse(w, book)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	newBook := appendNewBook(book)
	writeJSONResponse(w, newBook)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	updatedBook := updateTheBook(book)

	if nil != updatedBook {
		writeJSONResponse(w, updatedBook)
	} else {
		writeErrorMsg(w, "Error updating book")
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	id := getIDParams(r)

	success := deleteTheBook(id)
	if success {
		writeJSONResponse(w, WebResponse{"Successfully deleted"})
	}
	writeJSONResponse(w, WebResponse{"book not found"})
}

func (app *App) registerApis() {
	println("__registerApis__")

	app.router = mux.NewRouter()

	app.router.HandleFunc("/api/books", getBooks).Methods("GET")
	app.router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	app.router.HandleFunc("/api/books", createBook).Methods("POST")
	app.router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	app.router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":80", app.router))
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
	appendBooks()

	var app App = App{}
	app.registerApis()
	// app.start()
}
