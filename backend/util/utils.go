package util

import (
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfUserExists(userID string) *bson.M {
	// query MongoDB for user exists, if it does then return the document
	return nil
}
