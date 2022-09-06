package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
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
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helper.APIResponse("failed to create account", http.StatusUnprocessableEntity, "error", errorMessages)
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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	// Get input user
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Login proccess start
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessages := gin.H{"errors": err.Error()}

		response := helper.APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token := "token"
	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("login success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}