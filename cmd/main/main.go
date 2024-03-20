package main

import (
	"log"
	"net/http"
	

	"github.com/gorilla/mux"
                                                  // Import the routes package
	_ "github.com/jinzhu/gorm/dialects/mysql"    // Import MySQL driver as an anonymous import
)

func main() {
	// Create a new router instance
	r := mux.NewRouter()

	// Register routes for the bookstore
	//routes.RegisterBookStoreRoutes(r)
	

	// Start the HTTP server and listen on localhost:9010
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
