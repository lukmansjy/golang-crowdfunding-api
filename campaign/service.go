package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"golang-crowdfunding-api/user"
	"time"
)

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignsById(CampaignID int) (Campaign, error)
	CreateCampaign(input InputCampaign) (Campaign, error)
	UpdateCampaign(campaignId int, input InputCampaign) (Campaign, error)
	InsertImage(input InputImage, fileLocation string, currentUser user.User) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		return s.repository.FindByUserId(UserID)
	}
	return s.repository.FindAll()
}

func (s *service) GetCampaignsById(CampaignID int) (Campaign, error) {
	campaign, err := s.repository.FindById(CampaignID)
	return campaign, err
}

func (s *service) CreateCampaign(input InputCampaign) (Campaign, error) {
	slugFormat := fmt.Sprintf("%s-%d", input.Name, time.Now().UnixMilli())
	slugString := slug.Make(slugFormat)
	campaign := Campaign{
		UserID:           input.User.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		Slug:             slugString,
	}

	saveCampaign, err := s.repository.Save(campaign)
	return saveCampaign, err
}

func (s *service) UpdateCampaign(campaignId int, input InputCampaign) (Campaign, error) {
	campaign, err := s.repository.FindById(campaignId)
	if err != nil {
		return campaign, err
	}

	if campaign.User.ID != input.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	return s.repository.Update(campaign)
}

func (s *service) InsertImage(input InputImage, fileLocation string, currentUser user.User) (CampaignImage, error) {
	campaign, err := s.repository.FindById(input.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.User.ID != currentUser.ID {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1
		_, err := s.repository.MarkAllImageNotPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}

	campaignImage := CampaignImage{
		CampaignID: input.CampaignID,
		FileName:   fileLocation,
		IsPrimary:  isPrimary,
	}

	newCampaignImage, err := s.repository.SaveImage(campaignImage)
	return newCampaignImage, err
}
