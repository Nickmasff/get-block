package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type AppType string

const WebApp = AppType("webapp")

type Config struct {
	AppType  AppType
	GetBlock GetBlock
}

type GetBlock struct {
	BaseUrl string `yaml:"base_url"`
	Key     string `yaml:"key"`
}

func (cnf *Config) ReadConfig(configPath string) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	log.Print(fmt.Sprintf("ENVIRONMENT=%s", env))

	filePath, filePathErr := filepath.Abs(fmt.Sprintf("%s/config.%s.yml", configPath, env))
	if filePathErr != nil {
		log.Fatal(filePathErr.Error())
	}
	configRaw, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	cnf.setDefaultParams()
	err = yaml.Unmarshal(configRaw, &cnf)

	if err != nil {
		log.Fatal(fmt.Sprintf("config parse error: %v", err))
	}
}
func (cnf *Config) setDefaultParams() {
	cnf.AppType = WebApp
}
