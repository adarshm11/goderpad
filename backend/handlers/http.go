package handlers

import (
	"net/http"

	"goderpad/db"
	"goderpad/models"
	"goderpad/services"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var CreateRoomRequest models.CreateRoomRequest
	if err := c.ShouldBindJSON(&CreateRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := db.GetUserByID(c.Request.Context(), CreateRoomRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	accessLevel := user.AccessLevel
	if accessLevel < 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient access level to create room"})
		return
	}

	roomId, err := services.CreateRoom(CreateRoomRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roomID": roomId})
}

func DeleteRoom(c *gin.Context) {
	var DeleteRoomRequest models.DeleteRoomRequest
	if err := c.ShouldBindJSON(&DeleteRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := db.GetUserByID(c.Request.Context(), DeleteRoomRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	accessLevel := user.AccessLevel
	if accessLevel < 2 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient access level to create room"})
		return
	}
	err = services.DeleteRoom(DeleteRoomRequest)
	if err != nil {
		if _, ok := err.(*models.PermissionError); ok {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func JoinRoom(c *gin.Context) {
	var JoinRoomRequest models.JoinRoomRequest
	if err := c.ShouldBindJSON(&JoinRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := db.GetUserByID(c.Request.Context(), JoinRoomRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	err = services.JoinRoom(*user, JoinRoomRequest.RoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User joined room successfully"})
}

func LeaveRoom(c *gin.Context) {
	var LeaveRoomRequest models.LeaveRoomRequest
	if err := c.ShouldBindJSON(&LeaveRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := db.GetUserByID(c.Request.Context(), LeaveRoomRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	err = services.LeaveRoom(*user, LeaveRoomRequest.RoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User left room successfully"})
}
