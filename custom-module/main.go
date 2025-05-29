package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"custom-module/controllers"
	"custom-module/middlewares"
	"custom-module/services"
)

const defaultSize = 5

func main() {
	router := gin.Default()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	router.GET("/users", middlewares.NewPaginationWithDefaultSize(defaultSize), userController.GetAll)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
