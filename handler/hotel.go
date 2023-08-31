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

type HotelHandler struct {
	HotelService service.HotelService
	authService service.AuthService
}
func NewHotelHandler(HotelService service.HotelService, authService service.AuthService) HotelHandler {
	return HotelHandler{HotelService, authService}
}
func (h HotelHandler) Save(c *gin.Context) {
	var input dto.HotelInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newHotel, err := h.HotelService.Save(input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Hotel store failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatHotelDetail(newHotel)
	response := helper.APIResponse("Hotel has been store", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h HotelHandler) Update(c *gin.Context) {
	hotelId := c.Param("hotelId")
	var input dto.HotelInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Feature store failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newHotel, err := h.HotelService.Update(input, hotelId)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Hotel Update failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatHotelDetail(newHotel)
	response := helper.APIResponse("Hotel has been Updated", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}
func (h HotelHandler) GetHotel(c *gin.Context) {

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
	
	newHotel, err := h.HotelService.GetHotel(limit, page)
	if err != nil {
		errorMessage := gin.H{"errors": helper.FormatValidatorError(err)}
		response := helper.APIResponse("Hotel List failed2", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := formatter.FormatHotel(newHotel)
	response := helper.APIResponse("Hotel has been showing", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}