package main

import (
	"math/rand"
	"strconv"
)

var books []Book

//Book is book object
type Book struct {
	ID     string  `json:id`
	Isbn   string  `json:isbn`
	Title  string  `json:title`
	Author *Author `json:author`
}

//author is book's author
type Author struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
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

func appendBooks() {
	for i := 0; i < 10; i++ {
		index := strconv.Itoa(i)
		b := mockBook("Book #"+index, "12345"+index, "Fajar", "Al Munawwar")
		books = append(books, b)
	}
	println("BOOK SIZE: ", len(books))
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

func appendNewBook(book Book) Book {
	book.ID = getRandomID()
	books = append(books, book)

	return book
}

func updateTheBook(book Book) *Book {
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
		return &books[indexToReplace]
	}
	return nil

}

func deleteTheBook(id string) bool {
	for index, item := range books {
		if item.ID == id {
			books = append(books[:index], books[index+1:]...)

			return true
		}
	}
	return false
}
