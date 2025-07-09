package controller

import (
	"github.com/gin-gonic/gin"
	"golang-homework/task/taskFour/model"
	"golang-homework/task/taskFour/response"
	"golang-homework/task/taskFour/service"
	"net/http"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Register(ctx *gin.Context) {
	var user model.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		panic(response.Error(http.StatusBadRequest, "Invalid request body", nil))
	}
	if err := c.userService.RegisterUser(&user); err != nil {
		panic(response.Error(http.StatusInternalServerError, "Internal server error", nil))
	}
	response.ReturnSuccess(ctx, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var user model.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		panic(response.Error(http.StatusBadRequest, "Invalid request body", nil))
	}
	token, err := c.userService.Login(user.Username, user.Password)
	if err != nil {
		panic(response.Error(http.StatusInternalServerError, "Internal server error", nil))
	}
	user.Token = token
	response.ReturnSuccess(ctx, user)
}
