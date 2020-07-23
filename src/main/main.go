package main

import (
	"encoding/json"
	"log"
	"math/rand"
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

func writeResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

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

	id := getIDParams(r)

	println("get Books BY ID:", id)

	book := getBookByID(id)

	if nil == book {
		//	w.WriteHeader(404)
		writeResponseHeaders(w)
		writeErrorMsg(w, "Not Found")
	} else {
		writeResponseHeaders(w)
		writeJSONResponse(w, book)
	}
}

//WebResponse is response object
type WebResponse struct {
	Message string `json:message`
}

func writeErrorMsg(w http.ResponseWriter, msg string) {
	writeJSONResponse(w, WebResponse{msg})
}

func getBookByID(id string) *Book {
	for i := 0; i < len(books); i++ {
		b := books[i]
		if b.ID == id {
			return &b
		}
	}
	return nil
}

func writeJSONResponse(w http.ResponseWriter, obj interface{}) {
	json.NewEncoder(w).Encode(obj)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = getRandomID()
	books = append(books, book)
	writeJSONResponse(w, book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	println("will update book")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	indexToReplace := -1
loop:
	for index, item := range books {
		if item.ID == book.ID {
			indexToReplace = index
			break loop
		}
	}
	if indexToReplace >= 0 {
		books[indexToReplace] = book
		writeJSONResponse(w, books[indexToReplace])
	} else {
		writeErrorMsg(w, "Error updating book")
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	writeResponseHeaders(w)
	id := getIDParams(r)
	println("will delete by id: ", id)
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)
			writeJSONResponse(w, WebResponse{"Successfully deleted"})
			return
		}
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

func appendBooks() {
	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		b := mockBook("Book #"+index, "12345"+index, "Fajar", "AM")
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

func mockBook(title, isbn, authorName, authorLastName string) Book {
	book := Book{
		ID:     getRandomID(),
		Isbn:   isbn,
		Title:  title,
		Author: &Author{FirstName: authorName, LastName: authorLastName},
	}
	return book
}

func getRandomID() string {
	res := rand.Intn(10000000) + 10000000
	return strconv.Itoa(res)
}
