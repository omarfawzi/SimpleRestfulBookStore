package books

import (
	"Bookstore/controllers"
	"github.com/labstack/echo/v4"
)

type BookRouter struct {
	bookController *controllers.BookController
}

func New() *BookRouter {
	return &BookRouter{
		bookController: controllers.New(),
	}
}

func (router *BookRouter) RegisterRoutes(e *echo.Echo) {
	e.GET("/books", router.bookController.GetAll)
	e.GET("/books/:id", router.bookController.GetOne)
	e.DELETE("/books/:id", router.bookController.DeleteOne)
	e.POST("/books", router.bookController.Create)
	e.PUT("/books/:id", router.bookController.Update)
}
