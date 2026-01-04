package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Room struct {
	RoomName string
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	rooms := make(map[string]*Room)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		room, exists := rooms[roomId]
		if !exists {
			c.JSON(404, gin.H{
				"success": false,
				"error":   "Room not found",
			})
			return
		}
		c.JSON(200, gin.H{
			"success": true,
			"data": gin.H{
				"roomName": room.RoomName,
			},
		})
	})

	type CreateRoomRequest struct {
		Name     string `json:"name"`
		RoomName string `json:"roomName,omitempty"`
	}

	router.POST("/createRoom", func(c *gin.Context) {
		var req CreateRoomRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Invalid request",
			})
			return
		}
		roomId := uuid.New().String()

		log.Printf("Creating room for user: %s with name: %s and roomId: %s", req.Name, req.RoomName, roomId)
		rooms[roomId] = &Room{
			RoomName: req.RoomName,
		}
		c.JSON(200, gin.H{
			"success": true,
			"data": gin.H{
				"roomId": roomId,
			},
		})
	})

	type JoinRoomRequest struct {
		Name   string `json:"name"`
		RoomId string `json:"roomId"`
	}

	router.POST("/joinRoom", func(c *gin.Context) {
		var req JoinRoomRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"error":   "Invalid request",
			})
			return
		}
		log.Printf("User: %s joining roomId: %s", req.Name, req.RoomId)
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
