package transaction

import (
	"bwastartup/campaign"
	"bwastartup/payment"
	"errors"
	"fmt"
	"time"
)

type Service interface {
	GetCampaignTransactions(input CampaignTransactionsInput) ([]Transaction, error)
	GetUserTransactions(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository 		   Repository
	campaignRepository campaign.Repository
	paymentService 	   payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetCampaignTransactions(input CampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not the owner of the campaign")
	}

	transactions, err := s.repository.FindByCampaignID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}

func (s *service) GetUserTransactions(userID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	if err != nil {
		return []Transaction{}, err
	}

	return transactions, err
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.CampaignID)
	if err != nil {
		return Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return Transaction{}, errors.New("not the owner of the campaign")
	}

	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.Code = ""

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}
	
	paymentTransaction := payment.Transaction{}
	paymentTransaction.ID = newTransaction.ID
	paymentTransaction.Amount = newTransaction.Amount

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}
	newTransaction.PaymentURL = paymentURL

	year, _, _ := time.Now().Date()
	newTransaction.Code = fmt.Sprintf("BWA/%d/%d", year, newTransaction.ID)
	
	updatedTransaction, err := s.repository.Update(newTransaction)
	if err != nil {
		return updatedTransaction, err
	}

	return newTransaction, err
}
