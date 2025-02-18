package models

import (
	"book-store/pkg/config"

	"github.com/jinzhu/gorm"
)

// Book model

var db *gorm.DB
type Book struct {
	gorm.Model
	// ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"size:255;not null"`
	Author string `gorm:"size:255;not null"`
	Price  float64
}

func init (){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById (Id int64)(*Book,*gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

// UpdateBook updates an existing book
func (b *Book) UpdateBook() *gorm.DB {
	db := db.Save(&b)
	return db
}

// DeleteBook deletes a book by ID
func DeleteBook(Id int64) *gorm.DB {
	var book Book
	db := db.Where("ID = ?", Id).Delete(&book)
	return db
}
