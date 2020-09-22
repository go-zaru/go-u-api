package metaheaders

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// EnableRequestID ...
func EnableRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("x-request-id", uuid.New().String())
		c.Next()
	}
}
