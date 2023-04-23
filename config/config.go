package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const (
	DEV = "DEV"
	STG = "STG"
	PRO = "PRO"
)

type Config struct {
	Env struct {
		Dev Env `yaml:"dev"`
		Stg Env `yaml:"stg"`
		Pro Env `yaml:"pro"`
	} `yaml:"env"`
}

type Env struct {
	Db DatabaseConfig `yaml:"db"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

var globalConfig Config

func NewApplicationConfig() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occur")
		}
	}()
	filename, _ := filepath.Abs("./config/application.yml")

	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &globalConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Load global config successfully!!")
}

func GlobalConfig(env string) Env {
	switch env {
	case DEV:
		return globalConfig.Env.Dev
	case STG:
		return globalConfig.Env.Stg
	case PRO:
		return globalConfig.Env.Pro
	default:
		return globalConfig.Env.Dev
	}
}
