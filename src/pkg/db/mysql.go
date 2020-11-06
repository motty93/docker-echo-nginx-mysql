package db

import (
	"app/config"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

// Init is db session setting
func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := config.Config.DbUserName + ":" + config.Config.DbPassword + "@tcp(" + config.Config.DbHost + ":" + config.Config.DbPort + ")/" + config.Config.DbName + "?" + config.Config.DbParams
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error(err)
	}
}
