package repository

import (
	"errors"
	"fmt"
	"staycation/config"
	"staycation/domain"
)

type PaymentRepositoryContract interface {
	GetPaymentConfig() (domain.MidtransConfig, error)
}

type PaymentRepository struct {
	config config.PaymentConfigurations
}

func NewPaymentRepository(config config.PaymentConfigurations) PaymentRepository {
	return PaymentRepository{config}
}


func (r PaymentRepository) GetPaymentConfig() (domain.MidtransConfig, error) {
	var config domain.MidtransConfig

	config.APIEnv = r.config.APIEnv
	config.ClientKey = r.config.ClientKey
	config.ServerKey = r.config.ServerKey
	fmt.Println("ServerKEy")
	fmt.Println(r.config.ClientKey)
	if (config == domain.MidtransConfig{}) {
		return config, errors.New("failed to get config")
	}

	return config, nil
}