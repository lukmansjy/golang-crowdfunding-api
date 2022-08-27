package campaign

type FormatterCampaign struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) FormatterCampaign {
	imageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}
	campaignFormatter := FormatterCampaign{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		Slug:             campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         imageUrl,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []FormatterCampaign {
	//var campaignsResult []FormatterCampaign
	campaignsResult := []FormatterCampaign{}
	for _, campaign := range campaigns {
		formatCampaign := FormatCampaign(campaign)
		campaignsResult = append(campaignsResult, formatCampaign)
	}
	return campaignsResult
}
