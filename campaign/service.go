package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
	GetCampaignsById(CampaignID int) (Campaign, error)
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
