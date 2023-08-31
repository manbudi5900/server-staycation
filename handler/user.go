package handler

import (
	"fmt"
	"net/http"
	"staycation/domain"
	"staycation/dto"
	"staycation/formatter"
	"staycation/helper"
	"staycation/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
	authService service.AuthService
}
func NewUserHandler(userService service.UserService, authService service.AuthService) UserHandler {
	return UserHandler{userService, authService}
}
func (h UserHandler) RegisterUser(c *gin.Context) {
	var input dto.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h UserHandler) LoginUser(c *gin.Context) {
	
	var input dto.LoginUserInput


	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userLogin, err := h.userService.LoginUser(input)

	if err != nil {
		fmt.Println("error1")
		fmt.Println(err)
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	token, err := h.authService.GenerateToken(userLogin.ID)
	if err != nil {

		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.FormatUser(userLogin, token)
	response := helper.APIResponse("Account success login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h UserHandler) CekToken(c *gin.Context) {
	var input dto.CekTokenInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	_, err = h.authService.ValidateToken(input.Token)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func (h UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(domain.User)
	userID := currentUser.ID

	path := fmt.Sprintf("images/%s-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_upload": true}
	response := helper.APIResponse("Success to upload avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusAccepted, response)

}