package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	Http     HttpConf     `yaml:"http"`
	Log      LogConf      `yaml:"log"`
	Postgres PostgresConf `yaml:"postgres"`
}

type HttpConf struct {
	ListenPort int    `yaml:"listen_port"`
	RunMode    string `yaml:"run_mode"`
	AccessLog  string `yaml:"access_log"`
}

type LogConf struct {
	File   string `yaml:"file"`
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type PostgresConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db"`
	ShowSQL  string `yaml:"show_sql"`
}

func loadConfig(file *string) {
	f, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("load conf.json error: ", err)
	}
	err = yaml.Unmarshal(f, &Conf)
	if err != nil {
		fmt.Println("cannot unmarshal data:", err)
	}
	var tmp Config
	err = yaml.Unmarshal(f, &tmp)
	if err != nil {
		fmt.Println("cannot unmarshal data:", err)
	}
	G.Store("conf", tmp)
}

func initConfig(file *string) {
	loadConfig(file)
}
