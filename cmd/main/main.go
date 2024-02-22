package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/herondi/GO-BOOKSTORE/pkg/routes"
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Register the routes for the bookstore
	routes.RegisterBookStoreRoutes(r)

	// Handle the root URL with the router
	http.Handle("/", r)

	// Start the HTTP server and listen on localhost:9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

