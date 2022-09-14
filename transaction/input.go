package transaction

import "bwastartup/user"

type CampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}