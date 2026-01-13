package handlers

import (
	"log"

	"github.com/gin-gonic/gin"

	"goderpad/models"
	"goderpad/services"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func CreateRoom(c *gin.Context) {
	var req models.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   "Invalid request",
		})
		return
	}
	roomData, err := services.CreateRoom(req.UserID, req.Name, req.RoomName)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Internal server error",
		})
		return
	}
	log.Printf("Creating room with name: %s by user: %s", req.RoomName, req.Name)
	c.JSON(200, roomData)
}

func JoinRoom(c *gin.Context) {
	var req models.JoinRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   "Invalid request",
		})
		return
	}
	roomData, err := services.JoinRoom(req.UserID, req.Name, req.RoomId)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Internal server error",
		})
		return
	}
	if !roomData["success"].(bool) {
		c.JSON(404, roomData)
		return
	}
	log.Printf("User: %s joining roomId: %s", req.Name, req.RoomId)
	c.JSON(200, roomData)
}
