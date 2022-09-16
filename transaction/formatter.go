package transaction

import (
	"bwastartup/campaign"
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{}
	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	var campaignTransactionFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		campaignTransactionFormatter = append(campaignTransactionFormatter, FormatCampaignTransaction(transaction))
	}

	return campaignTransactionFormatter
}

type CampaignFormatter struct {
	Name 	 string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID 		  int 				`json:"id"`
	Amount    int 				`json:"amount"`
	Status    string 			`json:"status"`
	CreatedAt time.Time 		`json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

func FormatCampaign(campaign campaign.Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ImageUrl = ""

	for _, campaignImage := range campaign.CampaignImages {
		if campaignImage.IsPrimary {
			campaignFormatter.ImageUrl = campaignImage.FileName
		}
	}

	return campaignFormatter
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{}
	userTransactionFormatter.ID = transaction.ID
	userTransactionFormatter.Amount = transaction.Amount
	userTransactionFormatter.Status = transaction.Status
	userTransactionFormatter.CreatedAt = transaction.CreatedAt
	userTransactionFormatter.Campaign = FormatCampaign(transaction.Campaign)

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	var userTransactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		userTransactionsFormatter = append(userTransactionsFormatter, FormatUserTransaction(transaction))
	}

	return userTransactionsFormatter
}