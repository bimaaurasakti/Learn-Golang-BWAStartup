package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdateddAt time.Time
	User 	   user.User
}