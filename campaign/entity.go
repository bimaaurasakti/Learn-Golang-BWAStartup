package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID             	 int
	UserID         	 int
	Name           	 string
	ShortDescription string
	Description      string
	GoalAmount   	 int
	CurrentAmount 	 int
	Perks            string
	BackerCount      int
	Slug         	 string
	CreatedAt      	 time.Time
	UpdatedAt 	   	 time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	ID             	 int
	CampaignID       int
	FileName      	 string
	IsPrimary		 bool
	CreatedAt      	 time.Time
	UpdatedAt 	   	 time.Time
}