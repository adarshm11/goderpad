package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"goderpad/models"
	"goderpad/services"
)

func CreateRoomHandler(c *gin.Context) {
	var req models.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	roomID, err := services.CreateRoom(req.UserID, req.Name, req.RoomName)
	if err != nil {
		if errors.Is(err, models.ErrRoomExists) || errors.Is(err, models.ErrRoomNil) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "data": map[string]any{
		"roomId": roomID,
	}})
}

func JoinRoomHandler(c *gin.Context) {
	var req models.JoinRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := services.JoinRoom(req.UserID, req.Name, req.RoomID)
	if err != nil {
		if errors.Is(err, models.ErrRoomNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": map[string]any{
			"roomName": response["roomName"],
			"document": response["document"],
			"users":    response["users"],
		},
	})
}

func GetRoomNameHandler(c *gin.Context) {
	roomID := c.Param("roomID")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room ID is required"})
		return
	}

	roomName, err := services.GetRoomName(roomID)
	if err != nil {
		if errors.Is(err, models.ErrRoomNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"data": map[string]any{
			"roomName": roomName,
		},
	})
}
