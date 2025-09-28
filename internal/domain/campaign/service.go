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
		return "", internalErrors.ErrInternalServer
	}

	return campaign.ID, nil
}

func (s *Service) Delete(id string) error {

	_, err := s.Repository.GetById(id)
	if err != nil {
		return err
	}

	err = s.Repository.Delete(id)
	if err != nil {
		return internalErrors.ErrInternalServer
	}

	return nil
}

func (s *Service) Cancel(id string) error {

	_, err := s.Repository.GetById(id)
	if err != nil {
		return err
	}

	err = s.Repository.Cancel(id)
	if err != nil {
		return internalErrors.ErrInternalServer
	}

	return nil
}
