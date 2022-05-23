package main

import (
	"Faceit/src/internal/handlers"
	"Faceit/src/internal/repositories"
	service "Faceit/src/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {
	var db *mgo.Database
	colletionName := "users"

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	r := gin.Default()
	app := r.Group("/")

	NonRelationalUserDBRepository := repositories.NewMongoDBRepository(colletionName, db)
	userService := service.NewUserService(NonRelationalUserDBRepository)

	handlers.NewHealthHandler(app)
	handlers.NewUserHandler(app, userService)

	err := godotenv.Load("../../env/variables.env")
	if err != nil {
		logger.Error().Msg("Variables file not found...")
	}

	err = r.Run(":8080")
	if err != nil {
		logger.Error().Msg("Error running the server on port 8080")
	}

	logger.Info().Msg("Stopping server")

}
