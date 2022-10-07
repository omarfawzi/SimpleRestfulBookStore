package bootstrap

import (
	"Bookstore/database"
	errorMiddleware "Bookstore/middlewares/errors"
	"Bookstore/routers/books"
	"github.com/labstack/echo/v4"
)

type Bootstrap struct {
	echo       *echo.Echo
	bookRouter *books.BookRouter
	dbManager  *database.Manager
}

func New() *Bootstrap {
	return &Bootstrap{
		echo:       echo.New(),
		bookRouter: books.New(),
		dbManager:  database.Singleton(),
	}
}

func (bootstrap *Bootstrap) Start() {
	defer bootstrap.dbManager.CloseConnection()
	bootstrap.echo.Use(errorMiddleware.Handle)
	bootstrap.bookRouter.RegisterRoutes(bootstrap.echo)
	bootstrap.echo.Logger.Fatal(bootstrap.echo.Start(":1323"))
}
