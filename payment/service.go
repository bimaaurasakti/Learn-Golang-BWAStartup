package payment

import (
	"bwastartup/user"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct {

}

func NewService() *service {
	return &service{}
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	var client snap.Client

	client.New(goDotEnvVariable("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
	}

	SnapTokenRespon, err := client.CreateTransaction(snapReq)
	if err != nil {
		return "", err
	}

	return SnapTokenRespon.RedirectURL, nil
}