package campaign

import (
	"GoMailSender/internal/contract"
	"GoMailSender/internal/internalErrors"

)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaign)

	if err != nil {
		return "", errors.ErrInternalServer
	}

	return campaign.ID, nil
}
