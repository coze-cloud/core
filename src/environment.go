package core

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

const envKey = "APP_ENV"

func UseEnvironment(logger *logrus.Logger) {
	production := false
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		production = true
		_ = os.Setenv(envKey, "production")
	}

	if !production {
		_ = os.Setenv(envKey, "development")
		if err := godotenv.Load(); err != nil {
			logger.Warn("Development environment file could not be located")
		}
	}

	logger.Infof("Using %s environment", os.Getenv(envKey))
}

func IsDevelopment() bool {
	return os.Getenv(envKey) == "development"
}
