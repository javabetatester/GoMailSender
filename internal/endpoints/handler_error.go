package endpoints

import (
	"GoMailSender/internal/internalErrors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EndpointFunc define a assinatura padrão para endpoints que retornam dados, status HTTP e erro
type EndpointFunc func(ctx *gin.Context) (interface{}, int, error)

// HandlerError é um middleware que centraliza o tratamento de erros para todos os endpoints
// Elimina a necessidade de tratar erros individualmente em cada endpoint
func HandlerError(endpointFunc EndpointFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		obj, status, err := endpointFunc(ctx)

		// Tratamento centralizado de erros
		if err != nil {
			var httpStatus int
			
			// Determina o status HTTP baseado no tipo de erro
			if errors.Is(err, internalErrors.ErrInternalServer) {
				httpStatus = http.StatusInternalServerError
			} else {
				httpStatus = http.StatusBadRequest
			}
			
			ctx.JSON(httpStatus, gin.H{"error": err.Error()})
			return
		}

		// Resposta de sucesso
		if obj != nil {
			ctx.JSON(status, obj)
		} else {
			ctx.Status(status)
		}
	})
}
