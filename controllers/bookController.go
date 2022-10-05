package controllers

import (
	"Bookstore/models"
	"Bookstore/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BookController struct {
	bookService *services.BookService
}

func New() *BookController {
	return &BookController{
		bookService: services.New(),
	}
}

func (ctrl *BookController) GetAll(context echo.Context) error {
	err, books := ctrl.bookService.GetAll()
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, books)
}

func (ctrl *BookController) GetOne(context echo.Context) error {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err, book := ctrl.bookService.GetOne(id)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, book)
}

func (ctrl *BookController) DeleteOne(context echo.Context) error {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = ctrl.bookService.DeleteOne(id)
	if err != nil {
		return err
	}
	return context.NoContent(http.StatusOK)
}

func (ctrl *BookController) Create(context echo.Context) error {
	book := new(models.Book)

	err := context.Bind(book)
	if err != nil {
		return err
	}

	err, book = ctrl.bookService.Create(book)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, book)
}

func (ctrl *BookController) Update(context echo.Context) error {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		return err
	}

	book := new(models.Book)

	err = context.Bind(book)
	if err != nil {
		return err
	}

	err, book = ctrl.bookService.Update(id, book)

	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, book)
}
