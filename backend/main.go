package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goderpad/handlers"
	"goderpad/services"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/ws", handlers.HandleWebSocket)
	router.POST("/createRoom", handlers.CreateRoom)
	router.POST("/deleteRoom", handlers.DeleteRoom)
	router.POST("/joinRoom", handlers.JoinRoom)
	router.POST("/leaveRoom", handlers.LeaveRoom)

	go services.RegisterUsers()
	go services.UnregisterUsers()
	go services.StartRoomExpiration()

	defer services.StopHub()

	router.Run(":8080")
}
