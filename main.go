package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "encoding/json"
	// "log"
	// "net/http"
	// "math/rand"
	// "strconv"
	// "log"
	// "net/http"
	// "fmt"
	// "github.com/gorilla/mux"
)

//book struct (basically the model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//author struct
type Author struct {
	FirstName string `json"firstname"`
	LastName  string `json"lastname"`
}

//every function that we create for route handler has to have these two parameters. this is similar
// to the callback in app.get('/', (req, res)) in express

//get all books @GET
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api hit")
}

//get single book @GET
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create a new book @POST
func createBook(w http.ResponseWriter, r *http.Request) {

}

// update an existing book @PUT
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete an existing book @DELETE
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//init router
	r := mux.NewRouter()

	// create route handlers which will establish endpoints

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
