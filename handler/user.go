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