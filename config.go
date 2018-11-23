package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Manage struct {
		Username string
		Password string
	}
	DB struct {
		//Username string //- Ease 2018.11.23 改为sqlite3数据库
		//Password string //- Ease 2018.11.23 改为sqlite3数据库
		//Port     int64  //- Ease 2018.11.23 改为sqlite3数据库
		//Host     string //- Ease 2018.11.23 改为sqlite3数据库
		DBName string `yaml:"dbname"`
	} `yaml:"db"`
	DeployHost string `yaml:"deploy_host"`
	RunAt      string `yaml:"run_at"`
}

var config Config

func (self *Config) parse(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fd, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fd, self)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return &config
}
