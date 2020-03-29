package utils

import (
	"fmt"
)

// Init 初始化工具
func Init() {
	initArgs()
	fmt.Println("parse args ok")
	initConfig(&conf)
	fmt.Println("load config ok")
}
