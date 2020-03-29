package utils

import "sync"

var G sync.Map

func GetConf() Config {
	value, _ := G.Load("conf")
	return value.(Config)
}
