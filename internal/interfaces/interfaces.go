package interfaces

import "github.com/labstack/echo/v4"

type Route interface {
	RegisterRoute(group *echo.Group)
}
