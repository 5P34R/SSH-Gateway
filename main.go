package main

import (
	"ssh-gateway/gateway"
	"ssh-gateway/server"
)

func main() {
	server.Start(":8000")
	gateway.Start()
}
