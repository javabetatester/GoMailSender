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
	id, err := h.CampaignService.Create(request)
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
	campaigns, err := h.CampaignService.Repository.Get()
	if err != nil {
		return nil, 0, err // Status será determinado pelo tipo de erro
	}

	// Retorna sucesso
	return gin.H{"campaigns": campaigns}, http.StatusCreated, nil
}

// CampaignGetById busca a campanha pelo ID
// Tratamento de erros é delegado para o HandlerError middleware
func (h *Handler) CampaignGetById(c *gin.Context) (interface{}, int, error) {
	id := c.Param("id")
	campaign, err := h.CampaignService.Repository.GetById(id)
	if err != nil {
		return nil, 0, err // Status será determinado pelo tipo de erro
	}
	return campaign, http.StatusOK, nil
}
