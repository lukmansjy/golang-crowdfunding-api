package campaign

import "strings"

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

type FormatterCampaignDetail struct {
	ID               int                      `json:"id"`
	UserID           int                      `json:"user_id"`
	Name             string                   `json:"name"`
	Slug             string                   `json:"slug"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageUrl         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	Perks            []string                 `json:"perks"`
	User             FormatterUserCampaign    `json:"user"`
	Images           []FormatterImageCampaign `json:"images"`
}

type FormatterUserCampaign struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type FormatterImageCampaign struct {
	Name      string `json:"name"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatterImagesCampaign(images []CampaignImage) []FormatterImageCampaign {
	formatterImgCampaigns := []FormatterImageCampaign{}
	for _, image := range images {
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		formatterImgCampaign := FormatterImageCampaign{
			Name:      image.FileName,
			IsPrimary: isPrimary,
		}
		formatterImgCampaigns = append(formatterImgCampaigns, formatterImgCampaign)
	}
	return formatterImgCampaigns
}

func FormatCampaignDetail(campaign Campaign) FormatterCampaignDetail {
	imageUrl := ""
	if len(campaign.CampaignImages) > 0 {
		imageUrl = campaign.CampaignImages[0].FileName
	}

	formatterCampaignDetail := FormatterCampaignDetail{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		Slug:             campaign.Slug,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageUrl:         imageUrl,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Perks:            strings.Split(campaign.Perks, "|"),
	}

	formatterCampaignDetail.User = FormatterUserCampaign{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.AvatarFileName,
	}

	formatterCampaignDetail.Images = FormatterImagesCampaign(campaign.CampaignImages)

	return formatterCampaignDetail
}
