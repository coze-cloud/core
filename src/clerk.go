package core

import (
	"context"
	"github.com/cozy-hosting/clerk"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
)

var ClerkModule = fx.Provide(
	NewClerkConnection,
)

func NewClerkConnection(lifecycle fx.Lifecycle, logger *logrus.Logger) clerk.Connection {
	connectionString := os.Getenv("CLERK_CONNECTION_STRING")

	connection, err := clerk.NewMongoConnection(connectionString)
	if err != nil {
		logger.Fatal(err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			connection.Close(func(err error) {
				logger.Fatal(err)
			})
			return nil
		},
	})

	logger.Info("Successfully connected to MongoDB")
	return connection
}