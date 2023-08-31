package service

import (
	"errors"
	"fmt"
	"staycation/domain"
	"staycation/dto"
	"staycation/repository"

	"github.com/google/uuid"
)



type TransactionServiceContract  interface {
	CreateTransaction(input dto.TransactionInput) (domain.Transaction, error)
	GetTransactionsByUserID(userID string) ([]domain.Transaction, error)
	ProcessPayment(input dto.TransactionNotificationInput) error
	GetAllTransactions() ([]domain.Transaction, error)
  
  }
  
type TransactionService struct {
	TransactionRepository repository.TransactionRepository
	PaymentService PaymentService
	ProductRepository repository.ProductRepository
}

func NewTransactionService(TransactionRepository repository.TransactionRepository, PaymentService PaymentService, ProductRepository repository.ProductRepository) TransactionService{
	return TransactionService{TransactionRepository, PaymentService, ProductRepository}
}

func (s TransactionService) GetTransactionsByUserID(userID string) ([]domain.Transaction, error) {
	transactions, err := s.TransactionRepository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s TransactionService) CreateTransaction(input dto.TransactionInput) (domain.Transaction, error) {
	
	transaction := domain.Transaction{}
	transaction.Amount = input.Amount
	transaction.ProductID = input.ProductID
	transaction.UserID = input.User.ID
	transaction.Status = 0
	transaction.Code = ""
	transaction.ID = uuid.NewString()
	product, err := s.TransactionRepository.GetProductByID(input.ProductID)
	if err != nil {
		return transaction, errors.New("Product not found")
	}
	if product.IsBooking == 1 {
		return transaction, errors.New("This product has been booked")
	}

	newTransaction, err := s.TransactionRepository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := domain.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.PaymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	

	
	newTransaction.PaymentURL = string(paymentURL)
	newTransaction, err = s.TransactionRepository.Update(newTransaction)
	if err != nil {
		fmt.Println(err)
		return newTransaction, err
	}
	product.IsBooking = 1

	_, err = s.ProductRepository.Update(product)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
} 

func (s TransactionService) ProcessPayment(input dto.TransactionNotificationInput)(domain.Transaction, error) {
	var transaction_id = input.OrderID

	transaction, err := s.TransactionRepository.GetByID(transaction_id)
	if err != nil {
		return transaction ,err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "captured" && input.FraudStatus == "accept" {
		transaction.Status = 1
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = 1
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = -1
		product, err := s.TransactionRepository.GetProductByID(transaction.ProductID)
		if err != nil {
		}
		product.IsBooking = 0
		_, err = s.ProductRepository.Update(product)

	}

	updatedTransaction, err := s.TransactionRepository.Update(transaction)
	if err != nil {
		return transaction ,err
	}

	product, err := s.TransactionRepository.GetProductByID(updatedTransaction.ProductID)
	if err != nil {
		return transaction ,err
	}

	if updatedTransaction.Status == 1 {
		product.IsBooking =  1

		_, err := s.ProductRepository.Update(product)
		if err != nil {
			return transaction, err
		}
	}

	return transaction, nil
}

// func (s TransactionService) GetAllTransactions() ([]domain.Transaction, error) {
// 	transactions, err := s.TransactionRepository.FindAll()
// 	if err != nil {
// 		return transactions, err
// 	}

// 	return transactions, nil
// }