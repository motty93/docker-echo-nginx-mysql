package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type configList struct {
	DbUserName string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
	DbParams   string
}

// Config is exported variable configList
var Config configList

// Init function is config initialized
func Init() {
	// get current path
	p, _ := os.Getwd()
	fmt.Println(p)

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = configList{
		DbUserName: cfg.Section("mysql").Key("user_name").String(),
		DbPassword: cfg.Section("mysql").Key("password").String(),
		DbName:     cfg.Section("mysql").Key("db_name").String(),
		DbHost:     cfg.Section("mysql").Key("host").String(),
		DbPort:     cfg.Section("mysql").Key("port").String(),
		DbParams:   cfg.Section("mysql").Key("params").String(),
	}
}
