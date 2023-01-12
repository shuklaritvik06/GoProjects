package models

import (
	"github.com/shuklaritvik06/GoProjects/GO_MySQL_API/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	err := db.Create(&b).Error
	if err != nil {
		return nil, err
	}
	return b, err
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	d := db.Where("ID=?", id).Find(&book)
	if d.Error != nil {
		return nil, d
	}
	return &book, d
}

func GetAllBooks() (*[]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}

func DeleteBook(id int64) error {
	var book Book
	err := db.Where("ID=?", id).First(&book).Delete(&book).Error
	if err != nil {
		return err
	}
	return nil
}
