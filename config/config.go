package config

import (
	"github.com/spf13/viper"
	"log"
)

type configs struct {
	Application Application
	DataBase    DataBase
}

type Application struct {
	Host   string `yaml:"Application.host"`
	Port   string `yaml:"Application.port"`
	Prefix string `yaml:"Application.prefix"`
	Deploy string `yaml:"Application.deploy"`
}

type DataBase struct {
	Host        string `yaml:"DataBase.host"`
	Port        string `yaml:"DataBase.port"`
	User        string `yaml:"DataBase.user"`
	Password    string `yaml:"DataBase.password"`
	DBName      string `yaml:"DataBase.dbname"`
	TablePrefix string `yaml:"DataBase.tablePrefix"`
}

var Configs = &configs{}

func init() {
	viper.SetConfigFile("gcp-application.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("%v", err)
	}

	err = viper.Unmarshal(Configs)
	if err != nil {
		log.Fatalf("unable to decode into config struct, %v", err)
	}
}
