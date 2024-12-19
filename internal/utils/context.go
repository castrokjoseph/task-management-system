package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserIDFromContext(c *gin.Context) uuid.UUID {
	userID, exists := c.Get("user_id")
	if !exists {
		// This should never happen if auth middleware is used correctly
		panic("User ID not found in context")
	}
	return userID.(uuid.UUID)
}