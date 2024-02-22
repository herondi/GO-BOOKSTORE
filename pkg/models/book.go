package models

import (
	"github.com/akhil/GO-BOOKSTORE/pkg/config"
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
	db = config.GetDb()     // Get the database instance
	db.AutoMigrate(&Book{}) // Automatically migrate the Book struct to create the corresponding table in the database
}

func (b *Book) CreateBook() *Book {

	db.NewRecord(b) // Check if the record is new or not, If it's a new one then set the flag to true
	db.Create(&b)   // Call the create function on the object
	return b        // Return the created object back
}

func GetALLBooks() []Book { 
	var Books []Book                     // Slice of type Book
	db.Find(&Books) // Find all books in the database and store them in the variable 'Books'
	return Books  
}


func GetBookById(Id int64) (*Book, *gorm.DB){
var GetBook Book 
db:=db.Where("ID=?", Id).Find(&getBook)	// Find a book by its ID in the database
return &GetBook , db
}

func DeleteBook(ID int64){
	var book Book 
	db.Where("ID= ? ", ID).Delete(&book) // Delete a book by its ID from the database
	return book
	
}