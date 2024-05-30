package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins for development
    },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("WebSocket Upgrade:", err)
        return
    }
    defer conn.Close()

    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read Error:", err)
            break
        }
        // log.Printf("Received: %s", message)
        fmt.Println(string(message))
        // saveMessage(string(message))
    }
}


func setupRoutes() *gin.Engine {
    r := gin.Default()

    r.GET("/myApi2", func(c *gin.Context) {
        c.String(http.StatusOK, "Simple Server")
    })

    r.GET("/myApi", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome")
    })

    r.GET("/ws", func(c *gin.Context) {
        handleConnections(c.Writer, c.Request)
    })

    return r
}

func main() {
    // initDB()
    fmt.Println("Chat App v0.0.1")

    r := setupRoutes()

    // Root route to handle the root URL
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome to the Chat App Backend")
    })

    r.Run(":8080") // Listen and serve on port 8080 using the Gin router
}
