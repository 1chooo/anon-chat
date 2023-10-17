package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer conn.Close()

    scanner := bufio.NewScanner(conn)
    go func() {
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    }()

    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter message: ")
        text, _ := reader.ReadString('\n')
        conn.Write([]byte(text))
    }
}
