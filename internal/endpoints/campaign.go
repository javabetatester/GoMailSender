package endpoints

import (
	"GoMailSender/internal/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CampaignPost cria uma nova campanha
// Tratamento de erros é delegado para o HandlerError middleware
func (h *Handler) CampaignPost(c *gin.Context) (interface{}, int, error) {
	var request contract.NewCampaign

	// Bind do JSON - erro será tratado pelo HandlerError
	if err := c.ShouldBindJSON(&request); err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Criar campanha - erro será tratado pelo HandlerError
	id, err := h.CampaingService.Create(request)
	if err != nil {
		return nil, 0, err // Status será determinado pelo tipo de erro
	}

	// Retorna sucesso
	return gin.H{"id": id}, http.StatusCreated, nil
}

// CampaignGet busca todas as campanhas
// Tratamento de erros é delegado para o HandlerError middleware
func (h *Handler) CampaignGet(c *gin.Context) (interface{}, int, error) {
	// Buscar campanhas - erro será tratado pelo HandlerError
	campaigns, err := h.CampaingService.Repository.Get()
	if err != nil {
		return nil, 0, err // Status será determinado pelo tipo de erro
	}

	// Retorna sucesso
	return gin.H{"campaigns": campaigns}, http.StatusOK, nil
}
