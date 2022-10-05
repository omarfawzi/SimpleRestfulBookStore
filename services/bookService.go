package services

import (
	"Bookstore/database"
	"Bookstore/models"
	"database/sql"
	"github.com/go-playground/validator"
	"log"
)

type BookService struct {
	manager   *database.Manager
	validator *validator.Validate
}

func New() *BookService {
	return &BookService{
		manager:   database.Singleton(),
		validator: validator.New(),
	}
}

func (bookService *BookService) GetAll() (error, []models.Book) {
	query, err := bookService.manager.GetConnection().Query("Select * from books")
	if err != nil {
		return err, nil
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	var books []models.Book

	for query.Next() {
		book := new(models.Book)

		err = query.Scan(&book.Id, &book.Title, &book.Author)

		if err != nil {
			return err, nil
		}

		books = append(books, *book)
	}

	return nil, books
}

func (bookService *BookService) GetOne(id int64) (error, *models.Book) {
	book := new(models.Book)

	err := bookService.manager.GetConnection().QueryRow("select * from books where id = ?", id).Scan(&book.Id, &book.Title, &book.Author)

	if err != nil {
		return err, nil
	}

	return nil, book
}

func (bookService *BookService) DeleteOne(id int64) error {
	_, err := bookService.manager.GetConnection().Exec("DELETE from books where id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (bookService *BookService) Create(book *models.Book) (error, *models.Book) {
	err := bookService.validator.Struct(book)

	if err != nil {
		return err, nil
	}

	stmt, err := bookService.manager.GetConnection().Exec("INSERT INTO books (title, author) values (?, ?)", book.Title, book.Author)

	if err != nil {
		return err, nil
	}

	lastInsertedId, err := stmt.LastInsertId()

	if err != nil {
		return err, nil
	}

	err, book = bookService.GetOne(lastInsertedId)

	if err != nil {
		return err, nil
	}

	return nil, book
}

func (bookService *BookService) Update(id int64, book *models.Book) (error, *models.Book) {
	err := bookService.validator.Struct(book)

	if err != nil {
		return err, nil
	}

	_, err = bookService.manager.GetConnection().Exec("UPDATE books set title = ?, author = ? where id = ?", book.Title, book.Author, id)

	if err != nil {
		return err, nil
	}

	err, book = bookService.GetOne(id)

	if err != nil {
		return err, nil
	}

	return nil, book
}
