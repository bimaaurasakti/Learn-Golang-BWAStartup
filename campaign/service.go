package campaign


type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	var err error

	if userID == 0 {
		campaigns, err = s.repository.GetAll()
	} else {
		campaigns, err = s.repository.GetByUserID(userID)
	}

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
