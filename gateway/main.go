package main

import (
	controller "gateway/controllers"
	"gateway/db"
	"gateway/utils"
)

func main() {
	utils.Init()
	db.Init()
	controller.Start()
}
