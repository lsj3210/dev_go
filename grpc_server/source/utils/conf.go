package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	Server  ServerConf  `yaml:"server"`
	Statics StaticsConf `yaml:"statics"`
	Log     LogConf     `yaml:"log"`
	Mysql   MysqlConf   `yaml:"mysql"`
}

type ServerConf struct {
	GPort        int `yaml:"grpc_port"`
	GGatewayPort int `yaml:"grpc_gateway_port"`
}

type StaticsConf struct {
	SwaggerJsonPath string `yaml:"swagger_json_path"`
	SwaggerUiPath   string `yaml:"swagger_ui_path"`
}

type LogConf struct {
	File  string `yaml:"file"`
	Level string `yaml:"level"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
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
