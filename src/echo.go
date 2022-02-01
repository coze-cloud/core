package core

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

var EchoModule = fx.Provide(
	UseEcho,
)

func UseEcho(lifecycle fx.Lifecycle, logger *logrus.Logger) *echo.Echo {
	server := echo.New()

	server.Logger = &log.MyLogger{Logger: logger}
	server.Use(middleware.Logger())

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.Start(":1323"); err != nil && err != http.ErrServerClosed {
					logger.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(ctx); err != nil {
				logger.Fatal(err)
			}
			return nil
		},
	})

	return server
}
