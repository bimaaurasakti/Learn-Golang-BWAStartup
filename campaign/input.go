package campaign

import "bwastartup/user"

type CampaignDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type UploadCampaignImageInput struct {
	CampaignID int    `form:"campaign_id" binding:"required"`
	IsPrimary  string `form:"is_primary" binding:"required"`
	User       user.User
}