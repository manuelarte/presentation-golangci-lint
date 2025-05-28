package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"custom-module/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserController{service: service}
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users := c.service.GetAll()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
