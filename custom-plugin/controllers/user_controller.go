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
	page := ctx.GetInt("page")
	size := ctx.GetInt("size")

	users := c.service.GetAll(page, size)
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
