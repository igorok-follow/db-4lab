package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"database-service/config"
	"database-service/helpers/logger"
	"database-service/pkg/app"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "../../config/"

const defaultConfigName = "default_config.yaml"

var configName string

var configPath string

func init() {
	flag.StringVar(&configName, "config-name", defaultConfigName, "config name")
	flag.StringVar(&configPath, "config-path", defaultConfigPath, "config path")
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	defer func() {
		if r := recover(); r != nil {
			logger.Error("PANIC RECOVERY", r)
		}
	}()

	config, err := getConfig(configName)
	if err != nil {
		log.Printf("package main: config error \n%v", err)
	}

	app.Run(config)
}

func getConfig(name string) (*config.Config, error) {
	configPath = configPath + name

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config *config.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
