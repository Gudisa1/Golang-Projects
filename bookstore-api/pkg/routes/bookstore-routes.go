package routes

import (
	"book-store/pkg/controllers"

	"github.com/gorilla/mux"
)


var RegisterRoutes= func (r *mux.Router) {

	// Book routes
	r.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	
}
