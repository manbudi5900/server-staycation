package repository

import (
	"staycation/domain"

	"gorm.io/gorm"
)

type TransactionRepositoryContract interface {
	Save(product domain.Product)(domain.Product, error)
}
type TransactionRepository struct {
	
	DB *gorm.DB
  }
func (r TransactionRepository) DBReturn() (db *gorm.DB){
	return db
}
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{DB:db}
}

func (r TransactionRepository) Save(transaction domain.Transaction) (domain.Transaction, error) {
	
	err := r.DB.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
} 
func (r TransactionRepository) Update(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.DB.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r TransactionRepository) GetByID(ID string) (domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.DB.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
func (r *TransactionRepository) GetByUserID(userID string) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.DB.Preload("Product.ProductImage").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
func (r TransactionRepository) GetProductByID(ID string) (domain.Product, error) {
	var product domain.Product
	err := r.DB.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

