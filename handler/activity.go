package handler

import (
	"fmt"
	"net/http"
	"staycation/dto"
	"staycation/formatter"
	"staycation/helper"
	"staycation/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityHandler struct {
	ActivityService service.ActivityService
	authService service.AuthService
}
func NewActivityHandler(activityService service.ActivityService, authService service.AuthService) ActivityHandler {
	return ActivityHandler{activityService, authService}
}
func (h ActivityHandler) SaveActivity(c *gin.Context) {
	var input dto.ActivityInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Activity store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newActivity, err := h.ActivityService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Activity store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatActivity(newActivity)
	response := helper.APIResponse("Activity has been created", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h ActivityHandler) UpdateActivity(c *gin.Context) {
	activityId := c.Param("activityId")
	

	var input dto.ActivityInput




	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Activity store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newActivity, err := h.ActivityService.Update(input, activityId)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Activity store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatActivity(newActivity)
	response := helper.APIResponse("Activity has been created", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h ActivityHandler) GetActivity(c *gin.Context) {

	var page1 = c.DefaultQuery("page", "1")
	var limit1 = c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limit1)

    if err == nil {
        fmt.Println(limit1)
    }
	page, err := strconv.Atoi(page1)
	if err == nil {
        fmt.Println(limit1)
    }
	
	newActivity, err := h.ActivityService.GetActivity(limit, page)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Activity List failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Activity has been showing", http.StatusOK, "success", newActivity)
	c.JSON(http.StatusOK, response)
}
