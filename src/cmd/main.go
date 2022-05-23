package main

import (
	"Faceit/src/internal/handlers"
	"Faceit/src/internal/repositories"
	service "Faceit/src/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	var db *mgo.Database
	colletionName := "users"
	r := gin.Default()
	app := r.Group("/")

	NonRelationalUserDBRepository := repositories.NewMongoDBRepository(colletionName, db)
	userService := service.NewUserService(NonRelationalUserDBRepository)

	handlers.NewHealthHandler(app)
	handlers.NewUserHandler(app, userService)

	err := r.Run(":8080")
	if err != nil {
		//Log something with zerolog
	}

	//Log that server has stopped

}
