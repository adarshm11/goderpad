package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"

	"goderpad/handlers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/ping", handlers.Ping)

	router.POST("/createRoom", handlers.CreateRoom)
	router.POST("/joinRoom", handlers.JoinRoom)

	wsRouter := mux.NewRouter()
	wsRouter.HandleFunc("/ws/{roomId}", handlers.WebSocketHandler)

	router.Any("/ws/*any", gin.WrapH(wsRouter))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
