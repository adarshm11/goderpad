package handlers

import (
	"net/http"

	"goderpad/db"
	"goderpad/models"
	"goderpad/services"

	"github.com/gin-gonic/gin"
)

func authenticateUser(c *gin.Context, userID string, requiredAccessLevel int) (*models.User, error) {
	user, err := db.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		return nil, models.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	if user == nil {
		return nil, models.NewHTTPError(http.StatusUnauthorized, "User not authorized")
	}
	if user.AccessLevel < requiredAccessLevel {
		return nil, models.NewHTTPError(http.StatusForbidden, "Insufficient access level")
	}
	return user, nil
}

func CreateRoom(c *gin.Context) {
	var createRoomRequest models.CreateRoomRequest
	if err := c.ShouldBindJSON(&createRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := authenticateUser(c, createRoomRequest.UserID, models.OFFICER)
	if err != nil {
		if httpErr, ok := err.(*models.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	roomID, err := services.CreateRoom(createRoomRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"roomID": roomID})
}

func DeleteRoom(c *gin.Context) {
	var deleteRoomRequest models.DeleteRoomRequest
	if err := c.ShouldBindJSON(&deleteRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := authenticateUser(c, deleteRoomRequest.UserID, models.OFFICER)
	if err != nil {
		if httpErr, ok := err.(*models.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = services.DeleteRoom(deleteRoomRequest)
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
	var joinRoomRequest models.JoinRoomRequest
	if err := c.ShouldBindJSON(&joinRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authenticateUser(c, joinRoomRequest.UserID, models.MEMBER)
	if err != nil {
		if httpErr, ok := err.(*models.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = services.JoinRoom(user, joinRoomRequest.RoomID) // TODO: make this check if room exists, return 404 if not
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User joined room successfully"})
}

func LeaveRoom(c *gin.Context) {
	var leaveRoomRequest models.LeaveRoomRequest
	if err := c.ShouldBindJSON(&leaveRoomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := authenticateUser(c, leaveRoomRequest.UserID, models.MEMBER)
	if err != nil {
		if httpErr, ok := err.(*models.HTTPError); ok {
			c.JSON(httpErr.StatusCode, gin.H{"error": httpErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	err = services.LeaveRoom(user, leaveRoomRequest.RoomID) // TODO: make this check if room exists, return 404 if not
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User left room successfully"})
}
