package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"book-store/pkg/models"
	"book-store/pkg/utils"

	"github.com/gorilla/mux"
)

// CreateBook handles the creation of a new book

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books := models.GetAllBooks()
	if len(books) == 0 {
		response, _ := json.Marshal(map[string]string{"message": "No books found"})
		http.Error(w, string(response), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(books)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode books"}`, http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": "Invalid book ID"})
		http.Error(w, string(response), http.StatusBadRequest)
		return
	}

	book, _ := models.GetBookById(id)
	if book.ID == 0 {
		response, _ := json.Marshal(map[string]string{"error": "Book not found"})
		http.Error(w, string(response), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(book)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode book"}`, http.StatusInternalServerError)
		return
	}

	w.Write(response)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
    CreateBook :=&models.Book{}
    utils.ParseBody(r, &CreateBook)
	b:=CreateBook.CreateBook()
    response, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(response)

    // ...
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.ParseInt(params["id"], 10, 64)
  
	DeleteBook:=models.DeleteBook(id)
	res,_:=json.Marshal(DeleteBook)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
    
    

}
// GetAllBooks retrieves all books


// GetBookById retrieves a book by its ID

// UpdateBook updates an existing book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    book := &models.Book{}
    utils.ParseBody(r, book)
    vars := mux.Vars(r)
    bookId := vars["id"]
    id, err := strconv.ParseInt(bookId, 10, 64)
    if err != nil {
        http.Error(w, `{"error": "Invalid book ID"}`, http.StatusBadRequest)
        return
    }

    bookDetails, db := models.GetBookById(id)
    if db.Error != nil {
        http.Error(w, `{"error": "Failed to get book"}`, http.StatusInternalServerError)
        return
    }

    // Update the book details
    bookDetails.Title = book.Title
    bookDetails.Author = book.Author
    bookDetails.Price = book.Price

    db = db.Save(bookDetails)
    if db.Error != nil {
        http.Error(w, `{"error": "Failed to update book"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bookDetails)
}
