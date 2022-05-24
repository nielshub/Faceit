package main

import (
	"Faceit/src/internal/handlers"
	"Faceit/src/internal/repositories"
	service "Faceit/src/internal/services"
	"Faceit/src/log"
	"Faceit/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
)

func main() {
	var db *mgo.Database
	colletionName := "users"

	log.Init("debug")

	r := gin.Default()
	app := r.Group("/", middleware.LoggerMiddleware())

	NonRelationalUserDBRepository := repositories.NewMongoDBRepository(colletionName, db)
	userService := service.NewUserService(NonRelationalUserDBRepository)

	handlers.NewHealthHandler(app)
	handlers.NewUserHandler(app, userService)

	err := godotenv.Load("../../env/variables.env")
	if err != nil {
		log.Logger.Error().Msgf("Variables file not found...")
	}

	log.Logger.Info().Msgf("Starting server")
	err = r.Run(":8080")
	if err != nil {
		log.Logger.Error().Msgf("Error running the server on port 8080")
	}

	log.Logger.Info().Msgf("Stopping server")

}
