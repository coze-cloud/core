package core

import (
	"context"
	"github.com/cozy-hosting/messenger"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
)

var MessengerModule = fx.Provide(
	newMessenger,
)

func newMessenger(lifecycle fx.Lifecycle, logger *logrus.Logger) messenger.Messenger {
	connectionString := os.Getenv("MESSENGER_CONNECTION")

	msgr, err := messenger.NewRabbitMessenger(connectionString)
	if err != nil {
		logger.Fatal(err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			msgr.Close(func(err error) {
				logger.Fatal(err)
			})
			return nil
		},
	})

	logger.Info("Successfully connected to RabbitMQ")
	return msgr
}