package main

import (
	"GoMailSender/internal/domain/campaign"
	"GoMailSender/internal/endpoints"
	"GoMailSender/internal/infra/database"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	Service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaingService: Service,
	}

	g.POST("/gin-campaigns", 	endpoints.HandlerError(handler.CampaignPost))
	g.GET("/gin-campaigns", endpoints.HandlerError(handler.CampaignGet))

	g.Run(":3001")
}
