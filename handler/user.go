package handler

import (
	"golangsinau/helper"
	"golangsinau/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	// Binding the request body to the input variable.
	var input user.RegisterUserInput

	// Binding the request body to the input variable.
	err:= c.ShouldBindJSON(&input)
	if err!= nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

 // Checking if there is an error in the userService.RegisterUser(input) function. If there is an
 // error, it will return a bad request. If there is no error, it will return the user.
  Newuser, err:=h.userService.RegisterUser(input)
  if err!= nil {
	  response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(Newuser,"tokokokkokokokkokko")
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser,"tokokokkokokokkokko")
	response := helper.APIResponse("Successfully Loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context){
	var input user.CheckEmailInput
	
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error",errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isAvailable,
	}

	metaMessage := "Email has been registered"
	if isAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)


}