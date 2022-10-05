package routers

import "github.com/labstack/echo/v4"

type Router interface {
	RegisterRoutes(e *echo.Echo)
}
