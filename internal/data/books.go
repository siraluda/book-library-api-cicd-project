package data

import "database/sql"

type Book struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

type BookModel struct {
	DB *sql.DB
}

func (m BookModel) Get(Id int64) (*Book, error) {
	return &Book{ID: 1, Title: "Little Red Riding Hood", Author: "John Doe", Genre: "Fiction"}, nil
}

func (m BookModel) GetAll() (*[]Book, error) {
	return &[]Book{
		{ID: 1, Title: "Little Red Riding Hood", Author: "John Doe", Genre: "Fiction"},
		{ID: 2, Title: "Jack and the Beanstalk", Author: "Jane Doe", Genre: "Fiction"},
	}, nil
}

func (m BookModel) Insert(book *Book) error {
	// write sql query

	return nil
}

func (m BookModel) Delete(book *Book) error {
	// write sql query

	return nil
}
