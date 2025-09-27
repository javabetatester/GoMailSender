package endpoints

import "GoMailSender/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
