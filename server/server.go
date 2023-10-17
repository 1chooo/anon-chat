package main

import (
    "fmt"
    "net"
    "bufio"
    "sync"
)

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

    conn.Write([]byte("Enter your name: "))
    client.scanner.Scan()
    name := client.scanner.Text()
    client.name = name

    clientsMutex.Lock()
    clients[conn] = client
    clientsMutex.Unlock()

    broadcastMessage(name + " has joined the chat")

    for client.scanner.Scan() {
        message := client.scanner.Text()
        broadcastMessage(name + ": " + message)
    }

    clientsMutex.Lock()
    delete(clients, conn)
    clientsMutex.Unlock()

    broadcastMessage(name + " has left the chat")
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
    clients = make(map[net.Conn]Client)

    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Chat server started on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleClient(conn)
    }
}
