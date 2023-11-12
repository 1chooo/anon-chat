/*
Created Date: 2023/10/18
Author: @1chooo (Hugo ChunHo Lin)
Version: v0.0.1
*/

package main

import (
    "fmt"
    "net/http"

    "github.com/1chooo/socket-programming/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
    go websocket.Writer(ws)
    websocket.Reader(ws)
}

func setupRoutes() {
    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}
