package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Database struct {
		Name     string
		User     string
		Password string
		Host     string
	}

	Server struct {
		Port string
	}
}

var appConfig *AppConfig

func GetAppConfig() *AppConfig {

	if appConfig == nil {
		data, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}

		appConfig = &AppConfig{}
		err = yaml.Unmarshal(data, appConfig)
		if err != nil {
			panic(err)
		}
	}

	return appConfig
}
