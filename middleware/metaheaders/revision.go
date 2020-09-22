package metaheaders

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// EnableRevision handles setting meta data to help define requests etc
func EnableRevision(cfg *viper.Viper) gin.HandlerFunc {
	// Revision file contents will be only loaded once per process
	data := cfg.GetString("META_REVISION")

	// If we cant read file, just skip to the next request handler
	// This is pretty much a NOOP middlware :)
	if data == "" {
		// Make sure to log error so it could be spotted
		log.Println("revision middleware error: cannot retrive REVISION key")

		return func(c *gin.Context) {
			c.Next()
		}
	}

	// Clean up the value since it could contain line breaks
	revision := strings.TrimSpace(string(data))

	// Set out header value for each response
	return func(c *gin.Context) {
		c.Writer.Header().Set("x-revision", revision)
		c.Next()
	}
}
