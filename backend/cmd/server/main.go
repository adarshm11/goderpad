package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"goderpad/handlers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", handlers.Ping)

	router.GET("/room/:roomId", handlers.GetRoomInfo)
	router.POST("/createRoom", handlers.CreateRoom)
	router.POST("/joinRoom", handlers.JoinRoom)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
