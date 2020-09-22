package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*
  Microsoft IE :: Will not be supported from 17th August 2021
  So you can (and will) use lowercase headers

*/

// EnableCors enables cors for server
func EnableCors(cfg *viper.Viper) gin.HandlerFunc {

	// Consider how to use viper to pull config data

	defaultCfg := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"access-control-allow-headers", "content-type", "content-length", "accept-encoding", "x-csrf-token", "authorization", "accept", "origin", "cache-control", "x-requested-with"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	cors := cors.New(defaultCfg)
	// Set out header value for each response
	return cors
}

//  Essentially doing...
//  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//  c.Writer.Header().Set("Access-Control-Max-Age", "86400")
//  c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, UPDATE, OPTIONS")
//  c.Writer.Header().Set("access-control-allow-headers", "content-type", "content-length", "accept-encoding", "x-csrf-token", "authorization", "accept", "origin", "cache-control", "x-requested-with")
//  c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
//  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
