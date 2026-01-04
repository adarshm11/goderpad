package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	type CreateRoomRequest struct {
		Name   string `json:"name"`
		RoomId string `json:"roomId,omitempty"`
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
		if req.RoomId == "" {
			req.RoomId = uuid.New().String()
		}
		log.Printf("Creating room for user: %s with roomId: %s", req.Name, req.RoomId)
		c.JSON(200, gin.H{
			"success": true,
			"data": gin.H{
				"roomId": req.RoomId,
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
