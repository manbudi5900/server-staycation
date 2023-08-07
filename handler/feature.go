package handler

import (
	"fmt"
	"net/http"
	"staycation/dto"
	"staycation/formatter"
	"staycation/helper"
	"staycation/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	FeatureService service.FeatureService
	authService service.AuthService
}
func NewFeatureHandler(FeatureService service.FeatureService, authService service.AuthService) FeatureHandler {
	return FeatureHandler{FeatureService, authService}
}
func (h FeatureHandler) SaveFeature(c *gin.Context) {
	var input dto.FeatureInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newFeature, err := h.FeatureService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatFeature(newFeature)
	response := helper.APIResponse("Feature has been created", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h FeatureHandler) UpdateFeature(c *gin.Context) {
	featureId := c.Param("featureId")
	

	var input dto.FeatureInput




	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println("sasasa")
		fmt.Println(strings.Split(err.Error(), "Tag"))
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newFeature, err := h.FeatureService.Update(input, featureId)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatFeature(newFeature)
	response := helper.APIResponse("Feature has been created", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h FeatureHandler) GetFeature(c *gin.Context) {

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
	
	newFeature, err := h.FeatureService.GetFeature(limit, page)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature List failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Feature has been showing", http.StatusCreated, "success", newFeature)
	c.JSON(http.StatusOK, response)
}
