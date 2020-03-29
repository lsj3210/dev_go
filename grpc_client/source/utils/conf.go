package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	Server ServerConf `yaml:"server"`
}

type ServerConf struct {
	Addr string `yaml:"addr"`
	Func string `yaml:"func"`
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
