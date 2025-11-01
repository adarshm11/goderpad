package handlers

import (
	"goderpad/models"

	"goderpad/util"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	var CreateRoomRequest models.CreateRoomRequest
	if err := c.ShouldBindJSON(&CreateRoomRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := util.CheckIfUserExists(CreateRoomRequest.UserID)
	if user == nil {
		c.JSON(401, gin.H{"error": "User not authorized"})
		return
	}
	accessLevel := (*user)["accessLevel"].(int)
	if accessLevel < 2 {
		c.JSON(403, gin.H{"error": "Insufficient access level to create room"})
		return
	}
	// Generate unique Room ID and send into the register channel
}
