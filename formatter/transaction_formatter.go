package formatter

import (
	"staycation/domain"
	"time"
)


type ProductTransactionFormatter struct {
	ID        string       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type TransactionFormatter struct {
	ID         string    `json:"id"`
	ProductID string    `json:"product_id"`
	UserID     string    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     int `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

type UserTransactionFormatter struct {
	ID        string               `json:"id"`
	Amount    int               `json:"amount"`
	Status    int            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Product  ProductFormatter `json:"Product"`
}



func FormatProductTransaction(transaction domain.Transaction) ProductTransactionFormatter {
	formatter := ProductTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatProductTransactions(transactions []domain.Transaction) []ProductTransactionFormatter {
	if len(transactions) == 0 {
		return []ProductTransactionFormatter{}
	}

	var transactionsFormatter []ProductTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatProductTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

func FormatUserTransaction(transaction domain.Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	ProductFormatter := ProductFormatter{}
	ProductFormatter.Name = transaction.Product.Name
	ProductFormatter.ImageURL = ""

	if len(transaction.Product.ProductImage) > 0 {
		ProductFormatter.ImageURL = transaction.Product.ProductImage[0].Url
	}

	formatter.Product = ProductFormatter

	return formatter
}

func FormatUsersTransaction(transactions []domain.Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
func FormatTransaction(transaction domain.Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.ProductID = transaction.ProductID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}