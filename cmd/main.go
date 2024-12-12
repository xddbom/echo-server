package main

import (
	"ws-echo-server/server"
	"ws-echo-server/utils"
)

func main() {
	utils.InitLogger()
	utils.Log("Running server...")
	server.Start(":8080")
}
