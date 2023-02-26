package handlers

import (
	"fmt"
	"golang-gin-boilerplate/helper"
	"golang-gin-boilerplate/middlewares/authMiddleware"
	"golang-gin-boilerplate/services/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService authMiddleware.Service
}

func CreateUserHandler(userService user.Service, authService authMiddleware.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get input register struct
	input := &user.RegisterUserInput{}

	//bind the request body to input
	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Register account failed",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	newUser, err := h.userService.RegisterUser(*input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Register account failed",
			http.StatusBadRequest,
			"failed",
			err, nil)
		return
	}

	helper.SendResponse(
		c,
		"Your account successfully registered!",
		http.StatusOK,
		"success",
		user.FormatRegisterResponse(newUser))

}

func (h *userHandler) Login(c *gin.Context) {

	// init temp
	input := &user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)

	//if validation error
	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Login failed",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}
	loggedinUser, err := h.userService.Login(*input)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Login failed",
			http.StatusNotFound,
			"failed",
			err, nil)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.Id)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Login failed",
			http.StatusInternalServerError,
			"failed",
			err, nil)
		return
	}
	response := user.FormatLoginResponse(loggedinUser, token)
	helper.SendResponse(
		c,
		"You're successfully loggedin!",
		http.StatusOK,
		"success",
		response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	input := &user.CheckEmailInput{}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		helper.SendValidationErrorResponse(
			c,
			"Failed to check",
			http.StatusUnprocessableEntity,
			"failed",
			err)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(*input)

	responseData := gin.H{
		"is_available": isEmailAvailable,
		"email":        input.Email,
	}

	metaMessage := "Email is registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	helper.SendResponse(c, metaMessage, http.StatusOK, "success", responseData)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// get file from request with key avatar
	file, err := c.FormFile("avatar")

	response := gin.H{"is_uploaded": false}

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"failed",
			err, response)
		return
	}
	// get user from context data
	user := c.MustGet("loggedInUser").(user.User)
	userId := user.Id
	// create file path and filename
	path := fmt.Sprintf("images/avatar-%d-%s", time.Now().Unix(), file.Filename)

	// save uploaded file to filepath with filename
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"failed",
			err, response)
		return
	}

	// update user avatar path with id = userId in database with the path
	_, err = h.userService.SaveAvatar(userId, path)

	if err != nil {
		helper.SendErrorResponse(
			c,
			"Failed to upload avatar image",
			http.StatusBadRequest,
			"Failed",
			err, response)
		return
	}

	response = gin.H{"is_uploaded": true}

	helper.SendResponse(c, "Avatar successfully loaded", http.StatusOK, "success", response)
}
