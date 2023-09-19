package main

import (
	"ssh-gateway/server"
)

func main() {
	server.Start(":8000")
	
}
