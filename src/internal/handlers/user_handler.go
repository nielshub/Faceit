package handlers

import (
	"Faceit/src/internal/model"
	"Faceit/src/internal/ports"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	router      *gin.RouterGroup
	userService ports.UserService
}

func NewUserHandler(app *gin.RouterGroup, userService ports.UserService) UserHandler {
	userAPI := UserHandler{userService: userService}

	userRooter := app.Group("/user")
	userRooter.POST("/create", userAPI.userCreate)
	userRooter.PUT("/update", userAPI.userUpdate)
	userRooter.POST("/delete", userAPI.userDelete)
	userRooter.GET("/get", userAPI.getUsers)

	userAPI.router = userRooter

	return userAPI
}

func (uh *UserHandler) userCreate(c *gin.Context) {
	var newUser model.User
	newUserBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not read the body"})
		return
	}
	err = json.Unmarshal(newUserBody, &newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not unmarshal the json body"})
		return
	}

	newUser.InitializeTime()

	createdUser, err := uh.userService.CreateUser(c.Request.Context(), newUser)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
		return
	}
	response := "User has been created properly. User ID: " + createdUser.ID

	c.JSON(http.StatusOK, response)

}

func (uh *UserHandler) userUpdate(c *gin.Context) {
	id := c.Param("fileId")
	var updatedUser model.User
	updatedUserBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not read the body"})
		return
	}
	err = json.Unmarshal(updatedUserBody, &updatedUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not unmarshal the json body"})
		return
	}

	finalUser, err := uh.userService.UpdateUser(c.Request.Context(), id, updatedUser)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error updating user"})
		return
	}
	response := "User:" + finalUser.ID + " has been updated properly."

	c.JSON(http.StatusOK, response)

}

func (uh *UserHandler) userDelete(c *gin.Context) {
	id := c.Param("fileId")

	err := uh.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error deleting user"})
		return
	}
	response := "User:" + id + " has been deleted properly."

	c.JSON(http.StatusOK, response)

}

func (uh *UserHandler) getUsers(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")

	users, err := uh.userService.GetUsers(c.Request.Context(), key, value)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error getting users"})
		return
	}
	response := users

	c.JSON(http.StatusOK, response)

}
