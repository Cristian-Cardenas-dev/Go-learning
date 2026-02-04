package routes

import (
	"api-gateway/cmd/api/handlers"
	application "api-gateway/pkg/application/status"

	"github.com/gin-gonic/gin"
)

func RegisterStatusRoutes(router *gin.Engine) {
	useCase := application.NewStatusUseCase()
	handler := handlers.NewStatusGetHandler(*useCase)
	router.GET("/status", handler.Execute)
}
