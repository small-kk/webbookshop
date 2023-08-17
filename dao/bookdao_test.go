package dao

import (
	"book_mall/model"
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()

	for _, book := range books {
		fmt.Println(book)
	}
}

func TestAddBook(t *testing.T) {
	book := model.Book{
		Title:   "数据结构",
		Author:  "王道",
		Price:   69,
		Sales:   30,
		Stock:   70,
		ImgPath: "static/img/default.jpg",
	}
	err := AddBook(&book)
	fmt.Println(err)
}

func TestDeleteBook(t *testing.T) {
	book := &model.Book{
		ID: 32,
	}
	err := DeleteBook(book)
	fmt.Println(err)
}

func TestGetBookByID(t *testing.T) {
	book, _ := GetBookByID(28)
	fmt.Println(book)
}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      28,
		Title:   "机器学习",
		Author:  "杉山将",
		Price:   30,
		Sales:   200,
		ImgPath: "static/img/default.jpg",
	}
	err := UpdatesBook(book)
	fmt.Println(err)
}

func TestGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("7")
	fmt.Println(page)
}

func TestGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("4", "30", "100")
	fmt.Println(page)
}
