/*
Created Date: 2023/10/18
Author: @1chooo (Hugo ChunHo Lin)
Version: v0.0.1
*/

package main

import "fmt"
import "net"
import "bufio"
import "sync"

type Client struct {
	conn    net.Conn
	name    string
	scanner *bufio.Scanner
}

var clients map[net.Conn]Client
var clientsMutex sync.RWMutex

func handleClient(conn net.Conn) {
	defer conn.Close()

	client := Client{
		conn:    conn,
		scanner: bufio.NewScanner(conn),
	}

	client.scanner.Scan()
	name := client.scanner.Text()
	client.name = name

	clientsMutex.Lock()
	clients[conn] = client
	clientsMutex.Unlock()

	broadcastMessage(name + " has joined the chat")
	fmt.Println(name + " has joined the chat")

	for client.scanner.Scan() {
		message := client.scanner.Text()
		broadcastMessage(name + ": " + message)
	}

	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()

	broadcastMessage(name + " has left the chat")
	fmt.Println(name + " has left the chat")
}

func broadcastMessage(message string) {
	clientsMutex.RLock()
	defer clientsMutex.RUnlock()

	for conn, client := range clients {
		if _, err := conn.Write([]byte(message + "\n")); err != nil {
			// Handle error
			fmt.Printf("Error broadcasting message to %s: %v\n", client.name, err)
		}
	}
}

func main() {
	service := "localhost:1200"
	clients = make(map[net.Conn]Client)

	listener, err := net.Listen("tcp", service)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server started on :1200")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleClient(conn)
	}
}
