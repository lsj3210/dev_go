package main

import (
	"grpc_server/source/handlers"
	"grpc_server/source/utils"
)

func main() {
	utils.Init()
	handlers.Init()
}
