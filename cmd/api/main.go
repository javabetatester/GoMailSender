package main

import (
	"GoMailSender/internal/contract"
	"GoMailSender/internal/domain/campaign"
	"GoMailSender/internal/infra/database"
	"GoMailSender/internal/internalErrors"
	"errors"
	"net/http"

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

	g.POST("/gin-campaigns", func(c *gin.Context) {
		var request contract.NewCampaign

		// Gin tem binding automático mais simples
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := Service.Create(request)

		if err != nil {
			// Resposta de erro Interno
			if errors.Is(err, internalErrors.ErrInternalServer) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				// Resposta de erro Bad Request
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			return
		}

		// Resposta de sucesso
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	g.Run(":3001")
	//http.ListenAndServe(":3000", r)
}
