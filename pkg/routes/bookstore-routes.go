package routes

import (
	"./controllers"
	"github.com/gorilla/mux"
	"github.com/herondi/GO-BOOKSTORE/pkg/controllers"
)

// RegisterBookStoreRoutes registers the routes for the BookStore API.
// It takes a router as input and adds the necessary route handlers.
func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}