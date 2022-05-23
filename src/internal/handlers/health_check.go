package handlers

import (
	"Faceit/src/internal/model"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func NewHealthHandler(app *gin.RouterGroup) {
	app.GET("/health", health)
}

func health(c *gin.Context) {
	response := model.HealthResponse{
		Status:  http.StatusOK,
		Message: "Users microservice is up and running",
		Version: os.Getenv("VERSION"),
		Stack:   os.Getenv("ENVIRONMENT"),
	}

	c.JSON(http.StatusOK, response)
}
