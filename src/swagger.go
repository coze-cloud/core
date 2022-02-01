package core

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func UseSwagger(server *echo.Echo) {
	server.GET("/swagger", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusPermanentRedirect, "/swagger/")
	})
	server.GET("/swagger/*", echoSwagger.WrapHandler)
}
