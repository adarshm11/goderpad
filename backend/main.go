package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"goderpad/handlers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Failed to set websocket upgrade:", err)
			return
		}
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}

			name := string(message)
			response := fmt.Sprintf("Hello %s", name)

			err = conn.WriteMessage(messageType, []byte(response))
			if err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	})

	router.POST("/createRoom", handlers.CreateRoom)
	router.POST("/deleteRoom", handlers.DeleteRoom)
	router.POST("/joinRoom", handlers.JoinRoom)
	router.POST("/leaveRoom", handlers.LeaveRoom)

	router.Run(":8080")
}
