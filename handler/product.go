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
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService service.ProductService
	authService service.AuthService
}
func NewProductHandler(productService service.ProductService, authService service.AuthService) ProductHandler {
	return ProductHandler{productService, authService}
}


func (h *ProductHandler) UploadImage(c *gin.Context)  {
	file, err := c.FormFile("image")
	folder := c.Param("folder")
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var now = time.Now()
	

	path := fmt.Sprintf("images/"+folder+"/%s-%s-%s-%s-%s-%s", now.Month(), strconv.Itoa(now.Day()), strconv.Itoa(now.Hour()), strconv.Itoa(now.Minute()),strconv.Itoa(now.Second()), file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {

		data := gin.H{"is_upload": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "failed", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	

	data := gin.H{"is_upload": true, "path" : path}
	response := helper.APIResponse("Success to upload avatar image", http.StatusOK, "success", data)
	c.JSON(http.StatusAccepted, response)
}
func (h *ProductHandler) Save(c *gin.Context) {
	var input dto.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Println("sasasa")
		fmt.Println(strings.Split(err.Error(), "Tag"))
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Product store failed1", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newProduct, err := h.productService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Product store failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatProduct(newProduct)
	response := helper.APIResponse("Product has been created", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, response)
}