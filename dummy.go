package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/crypto/ssh"
)

func handleSession(user string, wg *sync.WaitGroup) {
	defer wg.Done()

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password("password")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "localhost:2201", config)
	if err != nil {
		log.Printf("Failed to connect to SSH server for user %s: %v", user, err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Printf("Failed to create SSH session for user %s: %v", user, err)
		return
	}
	defer session.Close()

	// Set up the terminal to reasonable values
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		log.Printf("Failed to allocate TTY for user %s: %v", user, err)
		return
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Printf("Failed to create stdin pipe for user %s: %v", user, err)
		return
	}

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Printf("Failed to create stdout pipe for user %s: %v", user, err)
		return
	}

	err = session.Shell()
	if err != nil {
		log.Printf("Failed to start shell for user %s: %v", user, err)
		return
	}

	// Create a channel to handle Ctrl+C
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		for {
			select {
			case <-sig:
				fmt.Fprint(stdin, "\x03") // Send Ctrl+C (ETX) to the shell
			}
		}
	}()

	// Monitor the session for closure
	go func() {
		err := session.Wait()
		if err != nil {
			log.Printf("Session closed for user %s: %v", user, err)
		}
	}()

	buffer := make([]byte, 4096)
	for {
		n, err := stdout.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Shell session closed for user %s.\n", user)
			} else {
				log.Printf("Failed to read from stdout for user %s: %v", user, err)
			}
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	// Session has closed or user has exited, we can exit this goroutine.
}

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go handleSession("root", &wg)
// 	wg.Wait()
// }
