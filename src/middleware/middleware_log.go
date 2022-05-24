package middleware

import (
	"Faceit/src/log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Logger.Info().Msgf("Endpoint Hit: " + c.Request.URL.Path)
		c.Next()
	}
}
