package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware(c *echo.Echo) {
	c.Use(middleware.Logger())
}