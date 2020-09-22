// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zaru/go-u-api/middleware/config"
	"github.com/go-zaru/go-u-api/middleware/cors"
	"github.com/go-zaru/go-u-api/middleware/metaheaders"
)

type cfgType struct{}

var cf cfgType

func main() {
	r := gin.Default()

	cfg := config.ReadConfig("")
	// enable cors
	r.Use(cors.EnableCors(cfg))
	// r.Use(authorisation.EnableCognitoAuth(cfg))
	// set x-request-id header
	r.Use(metaheaders.EnableRequestID())
	// set x-revision header
	r.Use(metaheaders.EnableRevision(cfg))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
