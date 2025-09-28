package database

import (
	"GoMailSender/internal/domain/campaign"
	"fmt"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign)
	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (c *CampaignRepository) GetById(id string) (campaign.Campaign, error) {
	var campaignResult campaign.Campaign
	tx := c.Db.Where("id = ?", id).First(&campaignResult)
	fmt.Println("DEBUGADO")
	return campaignResult, tx.Error
}

func (c *CampaignRepository) Delete(id string) error {
	// Primeiro verifica se a campanha existe
	var campaignToDelete campaign.Campaign
	tx := c.Db.Where("id = ?", id).First(&campaignToDelete)
	if tx.Error != nil {
		return tx.Error
	}

	// Deleta os contatos relacionados primeiro
	c.Db.Where("campaign_id = ?", id).Delete(&campaign.Contact{})

	// Depois deleta a campanha
	tx = c.Db.Delete(&campaignToDelete)
	return tx.Error
}

func (c *CampaignRepository) Cancel(id string) error {
	// Primeiro verifica se a campanha existe
	var campaignToCancel campaign.Campaign
	tx := c.Db.Where("id = ?", id).First(&campaignToCancel)
	if tx.Error != nil {
		return tx.Error
	}

	if campaignToCancel.Status == campaign.Canceled {
		return fmt.Errorf("campaign %s is already canceled", id)
	}

	campaignToCancel.Status = campaign.Canceled

	// Depois deleta a campanha
	tx = c.Db.Save(&campaignToCancel)
	return tx.Error
}
