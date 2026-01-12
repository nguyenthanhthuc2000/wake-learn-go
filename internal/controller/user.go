package controller

import (
	"golang/internal/service"
	"golang/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidRequest)
		return
	}

	result, err := uc.userService.GetUserByID(uint(userIDUint))
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeNotFound)
		return
	}
	response.SuccessResponse(c, response.Success, result)
}
