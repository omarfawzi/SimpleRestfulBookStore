package errors

import (
	errorFactory "Bookstore/factories/errors"
	"github.com/labstack/echo/v4"
)

func Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {
			mappedError := errorFactory.Make(err)
			return echo.NewHTTPError(mappedError.Code, mappedError.Message)
		}

		return nil
	}
}
