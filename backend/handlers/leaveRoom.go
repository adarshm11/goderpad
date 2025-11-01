package handlers

import (
	"goderpad/models"
	"goderpad/util"

	"github.com/gin-gonic/gin"
)

func LeaveRoom(c *gin.Context) {
	var LeaveRoomRequest models.LeaveRoomRequest
	if err := c.ShouldBindJSON(&LeaveRoomRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := util.CheckIfUserExists(LeaveRoomRequest.UserID)
	if user == nil {
		c.JSON(401, gin.H{"error": "User not authorized"})
		return
	}
	// Pass LeaveRoomRequest into the leave channel
}
