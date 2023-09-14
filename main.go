package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"ssh-gateway/utils"
	"sync"
	"time"
)

func main() {

	containerAddresses := utils.GetContainerAddress()

	listener, err := net.Listen("tcp", ":22")
	if err != nil {
		log.Fatalf("Failed to listen on port 22: %v", err)
	}
	defer listener.Close()

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection: %v", err)
			continue
		}
		fmt.Println(conn.RemoteAddr())
		containerAddress := containerAddresses[rand.Intn(len(containerAddresses))]

		wg.Add(1)
		go func(conn net.Conn, containerAddress string) {
			defer wg.Done()
			forwardConnection(conn, containerAddress)
		}(conn, containerAddress)
	}

	wg.Wait()
}

func forwardConnection(clientConn net.Conn, containerAddress string) {
	containerConn, err := net.Dial("tcp", containerAddress)
	if err != nil {
		log.Printf("Failed to connect to container %s: %v", containerAddress, err)
		clientConn.Close()
		return
	}
	defer containerConn.Close()

	go func() {
		_, err := io.Copy(clientConn, containerConn)
		if err != nil {
			log.Printf("Copy from container to client failed: %v", err)
		}
		clientConn.Close()
		containerConn.Close()
	}()

	_, err = io.Copy(containerConn, clientConn)
	if err != nil {
		log.Printf("Copy from client to container failed: %v", err)
	}
	clientConn.Close()
	containerConn.Close()
}
