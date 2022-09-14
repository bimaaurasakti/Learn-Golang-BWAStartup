package transaction

import (
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
	CreatedAt  time.Time
	UpdateddAt time.Time
	User 	   user.User
}