package main

import (
	"GoMailSender/internal/domain/campaign"
	"GoMailSender/internal/endpoints"
	"GoMailSender/internal/infra/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
_, err := database.NewDB()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco:", err)
	}
	defer database.CloseDB()

	g := gin.Default()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: service,
	}

	g.POST("/gin-campaigns", endpoints.HandlerError(handler.CampaignPost))
	g.GET("/gin-campaigns", endpoints.HandlerError(handler.CampaignGet))
	g.GET("/gin-campaigns/:id", endpoints.HandlerError(handler.CampaignGetById))

	log.Println("🚀 Servidor rodando na porta :3001 com GORM")
	g.Run(":3001")
}
