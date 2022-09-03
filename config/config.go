package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Database"`
	Server struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
	} `yaml:"Server"`
	Keys struct {
		PpApi        string `yaml:"PpApi"`
		Database     string `yaml:"Database"`
		AuraDatabase string `yaml:"AuraDatabase"`
	} `yaml:"Keys"`
}

func ReadConfig(filename string) (Config, error) {
	var Conf Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return Conf, err
	}

	yamlParser := yaml.NewDecoder(configFile)
	err = yamlParser.Decode(&Conf)

	if err != nil {
		fmt.Println(err)
	}

	return Conf, err
}

func getConfig() {
	fmt.Println("getting configurations")
}
