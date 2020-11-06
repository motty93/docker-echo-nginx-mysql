package main

import (
	"app/config"
	"app/pkg/db"
	"app/pkg/route"

	"github.com/sirupsen/logrus"
)

func init() {
	config.Init()
	db.Init()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	router := route.Router()

	router.Logger.Fatal(router.Start(":8020"))
}
