package endpoints

import "GoMailSender/internal/domain/campaign"

type Handler struct {
	CampaingService campaign.Service
}
