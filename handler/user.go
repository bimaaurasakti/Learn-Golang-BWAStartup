package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	// Get input user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMessages := gin.H{"errors": errors}

		response := helper.APIResponse("failed to create account", http.StatusBadRequest, "error", errorMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create user in database
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("failed to create account", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token := "token"
	userFormat := user.FormatUser(newUser, token)
	response := helper.APIResponse("your account has been craeted", http.StatusOK, "success", userFormat)

	c.JSON(http.StatusOK, response)
}