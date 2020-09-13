package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct Model
type Book struct {
	ID		string 	`json:"id"`
	Isbn	string 	`json:"isbn"`
	Title	string 	`json:"title"`
	Author	*Author `json:"author"`
}

// Get all Books
func getBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
// Get single Book
func getBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Book{})
}
// Create a New Book
func createBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}
// Update single Book
func updateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	json.NewEncoder(writer).Encode(books)

}
// Delete single Book
func deleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
		}
		break
	}
	json.NewEncoder(writer).Encode(books)
}

// Author Struck
type Author struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

func main() {
	// Init Router
	router := mux.NewRouter()

	// TODO Mock Data
	books = append(
		books, Book{
			ID: "1",
			Isbn: "18736591853",
			Title: "Book One",
			Author: &Author{
				Firstname: "John",
				Lastname: "Doe",
			},
		},
	)
	books = append(
		books, Book{
			ID: "2",
			Isbn: "18736747384",
			Title: "Book Two",
			Author: &Author{
				Firstname: "Alex",
				Lastname: "Megos",
			},
		},
	)

	//  Route Handlers / Endpoints
	router.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/v1/books", createBook).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DEL")

	log.Fatal(http.ListenAndServe(":8000", router))
}
