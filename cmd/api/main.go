package main

import (
	"GoMailSender/internal/contract"
	"GoMailSender/internal/domain/campaign"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	// ========== CHI ROUTER (EXISTENTE) ==========
	r := chi.NewRouter()

	service := campaign.Service{}
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {

		var request contract.NewCampaign
		render.DecodeJSON(r.Body, &request)

		id, err := service.Create(request)

		if err != nil {
			http.Error(w, err.Error(), 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	// ========== GIN WEB FRAMEWORK (EXEMPLO) ==========
	// Criando uma instância do Gin em paralelo para comparação
	ginRouter := gin.Default()

	// Middleware equivalente no Gin
	ginRouter.Use(gin.Logger())
	ginRouter.Use(gin.Recovery())

	ginService := campaign.Service{}
	ginRouter.POST("/gin-campaigns", func(c *gin.Context) {
		var request contract.NewCampaign

		// Gin tem binding automático mais simples
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := ginService.Create(request)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Resposta de sucesso
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	ginRouter.Run(":3001")
	http.ListenAndServe(":3000", r)
}
