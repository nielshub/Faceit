package main

import (
	"Faceit/src/internal/handlers"
	"Faceit/src/internal/repositories"
	service "Faceit/src/internal/services"
	"Faceit/src/log"
	"Faceit/src/middleware"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
)

func main() {
	// Start logger and load environment variables
	log.Init("debug")
	err := godotenv.Load("/usr/local/bin/variables.env")
	if err != nil {
		log.Logger.Error().Msgf("Variables file not found... Error: %s", err)
		panic(err)
	}
	log.Logger.Info().Msgf("Environment variables loaded")

	// Connect to mongoDB in Docker
	dbURL := os.Getenv("DBURL")
	usersCollectionName := os.Getenv("USERSCOLLECTIONNAME")
	dataBaseName := os.Getenv("DATABASENAME")
	session, err := mgo.Dial(dbURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	log.Logger.Info().Msgf("Connected to users DB")

	// Create repositories / services / handlers and app
	db := mgo.Database{
		Session: session,
		Name:    dataBaseName,
	}

	r := gin.Default()
	app := r.Group("/", middleware.LoggerMiddleware())

	NonRelationalUserDBRepository := repositories.NewMongoDBRepository(usersCollectionName, &db)
	userService := service.NewUserService(NonRelationalUserDBRepository)

	// Time to wait for rabbitMQ starts with docker-compose
	time.Sleep(20 * time.Second)

	publisherService := service.NewPublisherConnection("userEvents", "")
	if err := publisherService.Connect(); err != nil {
		panic(err)
	}
	log.Logger.Info().Msgf("Connected to publisher service")
	defer publisherService.Conn.Close()
	defer publisherService.Channel.Close()

	handlers.NewHealthHandler(app)
	handlers.NewUserHandler(app, userService, publisherService)

	// Run server
	log.Logger.Info().Msgf("Starting server")
	err = r.Run(":8080")
	if err != nil {
		log.Logger.Error().Msgf("Error running the server on port 8080. Error: %s", err)
	}

	log.Logger.Info().Msgf("Stopping server")

}
