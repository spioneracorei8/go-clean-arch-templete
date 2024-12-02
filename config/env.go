package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	APP_LOGGER bool
	APP_PORT   string
)

func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		logrus.Errorln(err)
	}
	APP_LOGGER = true
	APP_PORT = os.Getenv("APP_PORT")
}
