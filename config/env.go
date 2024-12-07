package config

import (
	"go-clean-arch-templete/helper"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	APP_LOGGER bool
	APP_PORT   string
	// cors
	ALLOW_HEADERS     string
	ALLOW_ORIGINS     string
	ALLOW_METHODS     string
	ALLOW_CREDENTIALS bool
	// mongo db
	// MONGODB_CONNECTION_URI string
)

func init() {
	var err error
	if err = godotenv.Load(); err != nil {
		logrus.Errorln(err)
	}
	APP_LOGGER = true
	APP_PORT = helper.GetENV("APP_PORT", "")

	// // cors
	ALLOW_HEADERS = helper.GetENV("ALLOW_HEADERS", "")
	ALLOW_ORIGINS = helper.GetENV("ALLOW_ORIGINS", "")
	ALLOW_METHODS = helper.GetENV("ALLOW_METHODS", "")
	if ALLOW_CREDENTIALS, err = strconv.ParseBool(helper.GetENV("ALLOW_CREDENTIALS", "")); err != nil {
		logrus.Errorln(err.Error())
	}

	// // mongo db
	// MONGODB_CONNECTION_URI = helper.GetENV("MONGODB_CONNECTION_URI", "")

}
