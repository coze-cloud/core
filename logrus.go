package core

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"os"
	"time"
)

var LogrusModule = fx.Provide(
	UseLogrus,
)

func UseLogrus() *logrus.Logger {
	logger := logrus.New()

	logrus.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC822,
	})

	return logger
}
