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

type CategoryProductHandler struct {
	categoryService service.CategoryService
	authService service.AuthService
}
func NewCategoryProductHandler(categoryService service.CategoryService, authService service.AuthService) CategoryProductHandler {
	return CategoryProductHandler{categoryService, authService}
}
func (h CategoryProductHandler) SaveCategory(c *gin.Context) {
	var input dto.CategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println("sasasa")
		fmt.Println(strings.Split(err.Error(), "Tag"))
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Category store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newCategory, err := h.categoryService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Category store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatCategroy(newCategory)
	response := helper.APIResponse("Category has been created", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}

func (h CategoryProductHandler) UpdateCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")
	

	var input dto.CategoryInput




	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println("sasasa")
		fmt.Println(strings.Split(err.Error(), "Tag"))
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Category store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newCategory, err := h.categoryService.Update(input, categoryId)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Category store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatCategroy(newCategory)
	response := helper.APIResponse("Category has been updated", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h CategoryProductHandler) GetCategory(c *gin.Context) {

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
	
	newCategory, err := h.categoryService.GetCategory(limit, page)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Category List failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Category has been showing", http.StatusOK, "success", newCategory)
	c.JSON(http.StatusOK, response)
}
