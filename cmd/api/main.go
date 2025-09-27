package main

import (
	"GoMailSender/internal/domain/campaign"
	"GoMailSender/internal/endpoints"
	"GoMailSender/internal/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	// Middleware equivalente no Gin
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	Service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaingService: Service, // CHAMA O SERVICE ACIMA
	}

	g.POST("/gin-campaigns", handler.CampaignPost)

	g.Run(":3001")
}
