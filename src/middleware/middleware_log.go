package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
		logger.Info().Msg("Endpoint Hit: " + c.Request.URL.Path)
		c.Next()
	}
}
