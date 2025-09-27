package endpoints

import (
	"GoMailSender/internal/contract"
	"GoMailSender/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CampaignPost(c *gin.Context) {
	var request contract.NewCampaign

	// Gin tem binding automático mais simples
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.CampaingService.Create(request)

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
}
