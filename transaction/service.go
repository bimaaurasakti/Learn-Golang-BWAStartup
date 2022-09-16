package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetCampaignTransactions(input CampaignTransactionsInput) ([]Transaction, error)
	GetUserTransactions(userID int) ([]Transaction, error)
}

type service struct {
	repository 		   Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
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