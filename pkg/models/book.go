package models

import (
	"github.com/herondi/GO-BOOKSTORE/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`               // Name of the book
	Author      string `gorm:"column:author" json:"author"`           // Author of the book
	Publication string `gorm:"column:publication" json:"publication"` // Publication of the book
}

func init() {
	config.Connect()        // Connect to the database
	db = config.GetDB()     // Get the database instance (corrigido para GetDB)
	db.AutoMigrate(&Book{}) // Automatically migrate the Book struct to create the corresponding table in the database
}

// CreateBook creates a new book in the database and returns it.
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Check if the record is new, if it is, set the flag to true
	db.Create(&b)   // Call the create function on the object
	return b        // Return the created object back
}

// GetAllBooks returns all books in the database.
func GetAllBooks() []Book {
	var books []Book // Slice of books
	db.Find(&books)  // Find all books in the database and store them in the variable 'books'
	return books     // Return the books
}

// GetBookById returns a book with the provided ID.
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book // Initialize a Book variable to store the result
	db := db.Where("ID=?", Id).Find(&book) // Find a book by its ID in the database
	return &book, db // Return a pointer to the book and the database connection
}

// DeleteBook deletes a book with the provided ID.
func DeleteBook(ID int64) {
	var book Book // Initialize a Book variable to store the result
	db.Where("ID= ?", ID).Delete(&book) // Delete a book by its ID from the database
}
