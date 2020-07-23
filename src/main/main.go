package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

type Book struct {
	ID     string  `json:id`
	Isbn   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:author`
}

type Author struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	println("===getBooks===")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	println("===getBooks===")
	params := mux.Vars(r)
	id := params["id"]
	book := getBookById(id)
	json.NewEncoder(w).Encode(book)
}

func getBookById(id string) *Book {
	for i := 0; i < len(books); i++ {
		b := books[i]
		if b.ID == id {
			return &b
		}
	}
	return nil
}

func createBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

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

func appendBooks() {
	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		b := mockBook("Book #"+index, index, "12345"+index, "Fajar", "AM")
		books = append(books, b)
	}
	println("BOOK SIZE: ", len(books))
}

func main() {
	appendBooks()

	var app App = App{}
	app.registerApis()
	// app.start()
}

func mockBook(title, id, isbn, authorName, authorLastName string) Book {
	book := Book{
		ID:     id,
		Isbn:   isbn,
		Title:  title,
		Author: &Author{FirstName: authorName, LastName: authorLastName},
	}
	return book
}
