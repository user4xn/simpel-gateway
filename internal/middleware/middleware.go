package middleware

import (
	"net/http"
	"regexp"
	"simpel-gateway/pkg/util"

	"github.com/gin-gonic/gin"
)

func AuthBearerServer() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		if len(header) == 0 {
			response := util.APIResponse("Sorry, you didn't enter a valid bearer token", http.StatusUnauthorized, "failed", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		rep := regexp.MustCompile(`(Bearer)\s?`)
		bearerStr := rep.ReplaceAllString(header[0], "")
		bearerEnv := util.GetEnv("BEARER_SERVER", "fallback")
		if bearerStr != bearerEnv {
			response := util.APIResponse("Sorry, you didn't enter a valid bearer token", http.StatusUnauthorized, "failed", nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		c.Next()
	}
}
