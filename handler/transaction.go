package handler

import (
	"fmt"
	"log"
	"net/http"
	"staycation/domain"
	"staycation/dto"
	"staycation/formatter"
	"staycation/helper"
	"staycation/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	transactionService service.TransactionService
	authService service.AuthService
}
func NewTransactionHandler(transactionService service.TransactionService, authService service.AuthService) TransactionHandler {
	return TransactionHandler{transactionService, authService}
}

// func (h TransactionHandler) GetCampaignTransactions(c *gin.Context) {
// 	var input dto.GetCampaignTransactionsInput
// 	err := c.ShouldBindUri(&input)
// 	if err != nil {
// 		response := helper.APIResponse(
// 			"Error getting campaign's transactions",
// 			http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	transactions, err := h.service.GetTransactionsByCampaignID(input)
// 	if err != nil {
// 		response := helper.APIResponse(
// 			"Failed to get campaign's transactions",
// 			http.StatusBadRequest, "error", nil)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response := helper.APIResponse(
// 		"Campaign's transaction",
// 		http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
// 	c.JSON(http.StatusOK, response)
// }

func (h TransactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(domain.User)
	userID := currentUser.ID
	

	transactions, err := h.transactionService.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get user's transactions",
			http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User's transaction",
		http.StatusOK, "success", formatter.FormatUsersTransaction(transactions))
	c.JSON(http.StatusOK, response)
}

func (h TransactionHandler) CreateTransaction(c *gin.Context) {
	var input dto.TransactionInput
	txHandle := c.MustGet("db_trx").(*gorm.DB)

	fmt.Println(txHandle)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		
		txHandle.Rollback()
		errors := helper.FormatValidatorError(err)
		fmt.Println("rollback1")

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fmt.Println(input)

	currentUser := c.MustGet("currentUser").(domain.User)
	input.User = currentUser

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		txHandle.Rollback()
		errors := helper.FormatValidatorError(err)

		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Create transaction failed",
			http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	if err := txHandle.Commit().Error; err != nil {
		log.Print("trx commit error: ", err)
	}


	formatter := formatter.FormatTransaction(newTransaction)
	response := helper.APIResponse(
		"Success create transaction",
		http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h TransactionHandler) GetNotification(c *gin.Context) {
	var input dto.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatorError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to process notification",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.transactionService.ProcessPayment(input)
	if err != nil {
		errors := helper.FormatValidatorError((err))
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(
			"Failed to process notification",
			http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
