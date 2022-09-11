package campaign

import (
	"bwastartup/user"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormat := CampaignFormatter{}
	campaignFormat.ID = campaign.ID
	campaignFormat.UserID = campaign.UserID
	campaignFormat.Name = campaign.Name
	campaignFormat.ShortDescription = campaign.ShortDescription
	campaignFormat.GoalAmount = campaign.GoalAmount
	campaignFormat.CurrentAmount = campaign.CurrentAmount
	campaignFormat.Slug = campaign.Slug
	campaignFormat.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormat.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormat
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	if len(campaigns) == 0 {
		return campaignsFormatter
	}

	for _, campaign := range campaigns {
		campaignsFormatter = append(campaignsFormatter, FormatCampaign(campaign))
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	UserID           int                       `json:"user_id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	ImageUrl         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	Description      string                    `json:"description"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ", ") {
		perks = append(perks, perk)
	}
	campaignDetailFormatter.Perks = perks
	campaignDetailFormatter.User = FormatCampaignUserFormatter(campaign.User)
	campaignDetailFormatter.Images = FormatCampaignImagesFormatter(campaign.CampaignImages)

	return campaignDetailFormatter
}

func FormatCampaignUserFormatter(user user.User) CampaignUserFormatter {
	campaignUserFormater := CampaignUserFormatter{}
	campaignUserFormater.Name = user.Name
	campaignUserFormater.ImageUrl = user.AvatarFileName

	return campaignUserFormater
} 

func FormatCampaignImagesFormatter(campaignImages []CampaignImage) []CampaignImagesFormatter {
	var campaignImagesFormatterArray []CampaignImagesFormatter

	for _, images := range campaignImages {
		campaignImagesFormatter := CampaignImagesFormatter{}
		campaignImagesFormatter.ImageUrl = images.FileName
		campaignImagesFormatter.IsPrimary = false
		if images.IsPrimary == "1" {
			campaignImagesFormatter.IsPrimary = true
		}

		campaignImagesFormatterArray = append(campaignImagesFormatterArray, campaignImagesFormatter)
	}

	return campaignImagesFormatterArray
}