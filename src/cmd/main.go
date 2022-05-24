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
	log.Init("debug")
	err := godotenv.Load("./env/variables.env")
	if err != nil {
		log.Logger.Error().Msgf("Variables file not found... Error: %s", err)
		return
	}
	log.Logger.Info().Msgf("Environment variables loaded")

	//When deploying app in Docker
	dbURL := "mongodb"
	colletionName := "users"
	dataBaseName := "faceit"
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Logger.Error().Msgf("Error connecting to db. Error: %s", err)
		return
	}
	log.Logger.Info().Msgf("Connected to users DB")

	db := mgo.Database{
		Session: session,
		Name:    dataBaseName,
	}

	r := gin.Default()
	app := r.Group("/", middleware.LoggerMiddleware())

	NonRelationalUserDBRepository := repositories.NewMongoDBRepository(colletionName, &db)
	userService := service.NewUserService(NonRelationalUserDBRepository)

	handlers.NewHealthHandler(app)
	handlers.NewUserHandler(app, userService)

	log.Logger.Info().Msgf("Starting server")
	err = r.Run(":8080")
	if err != nil {
		log.Logger.Error().Msgf("Error running the server on port 8080. Error: %s", err)
	}

	log.Logger.Info().Msgf("Stopping server")

}
