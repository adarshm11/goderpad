package handlers

import (
	"goderpad/models"
	"goderpad/util"

	"github.com/gin-gonic/gin"
)

func DeleteRoom(c *gin.Context) {
	var DeleteRoomRequest models.DeleteRoomRequest
	if err := c.ShouldBindJSON(&DeleteRoomRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := util.CheckIfUserExists(DeleteRoomRequest.UserID)
	if user == nil {
		c.JSON(401, gin.H{"error": "User not authorized"})
		return
	}
	accessLevel := (*user)["accessLevel"].(int)
	if accessLevel < 2 {
		c.JSON(403, gin.H{"error": "Insufficient access level to create room"})
		return
	}
	// Pass DeleteRoomRequest into the unregister channel
}
