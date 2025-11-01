package handlers

import (
	"goderpad/models"
	"goderpad/util"

	"github.com/gin-gonic/gin"
)

func JoinRoom(c *gin.Context) {
	var JoinRoomRequest models.JoinRoomRequest
	if err := c.ShouldBindJSON(&JoinRoomRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := util.CheckIfUserExists(JoinRoomRequest.UserID)
	if user == nil {
		c.JSON(401, gin.H{"error": "User not authorized"})
		return
	}
	// Pass JoinRoomRequest into the join channel
}
