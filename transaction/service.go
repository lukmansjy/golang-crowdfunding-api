package transaction

type Service interface {
	GetTransactionsByCampaignID(campaignID int) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(campaignID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignID(campaignID)
	return transactions, err
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserID(userID)
	return transactions, err
}
