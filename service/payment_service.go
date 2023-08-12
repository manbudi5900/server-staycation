package service

import (
	"fmt"
	"staycation/domain"
	"staycation/repository"

	midtrans "github.com/veritrans/go-midtrans"
)

type PaymentServiceContract interface {
	GetPaymentURL(transaction domain.Transaction, user domain.User) (string, error)
}

type PaymentService struct {
	PaymentRepository repository.PaymentRepository
}

func NewPaymentService(repository repository.PaymentRepository) PaymentService {
	return PaymentService{repository}
}



func (s PaymentService) GetPaymentURL(transaction domain.Transaction, user domain.User) (string, error) {
	midclient := midtrans.NewClient()
	

	config, err := s.PaymentRepository.GetPaymentConfig()
	fmt.Println("12131121");
	fmt.Println(err);
	if err != nil {
		return "", err
	}
	fmt.Println("asaadsefrgthjy");

	midclient.ClientKey = config.ClientKey
	midclient.ServerKey = config.ServerKey
	env := config.APIEnv
	fmt.Println("env")
	fmt.Println(env)
	if env == "production" {
		midclient.APIEnvType = midtrans.Production
	}
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.ID,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}