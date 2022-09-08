package campaign

import "time"

type Campaign struct {
	ID             	 int
	UserID         	 int
	Name           	 string
	ShortDescription string
	Description      string
	GoalAmount   	 string
	CurrentAmount 	 string
	Perks            string
	BackerCount      string
	Slug         	 string
	CreatedAt      	 time.Time
	UpdatedAt 	   	 time.Time
}

type CampaignImage struct {
	ID             	 int
	UserID         	 int
	Name           	 string
	ShortDescription string
	Description      string
	GoalAmount   	 string
	CurrentAmount 	 string
	Perks            string
	BackerCount      string
	Slug         	 string
	CreatedAt      	 time.Time
	UpdatedAt 	   	 time.Time
}