package handlers

import (
	"net/http"

	"api-gateway/pkg/application/status"

	"github.com/gin-gonic/gin"
)

type StatusGetHandler struct {
	useCase application.StatusUseCase
}

func NewStatusGetHandler(uc application.StatusUseCase) *StatusGetHandler {
	return &StatusGetHandler{
		useCase: uc,
	}
}

func (ctrl *StatusGetHandler) Execute(ctx *gin.Context) {
	message := ctrl.useCase.Execute()
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
