package campaign

type CampaignFormatter struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl string `json:"image_url"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`
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