package handlers

import (
	"net/http"

	"api-gateway/cmd/api/core"
	application "api-gateway/pkg/application/user"

	"github.com/gin-gonic/gin"
)

type UserPostHandler struct {
	useCase application.UserCreator
}

func NewUserPostHandler(uc application.UserCreator) *UserPostHandler {
	return &UserPostHandler{
		useCase: uc,
	}
}

func (ctrl *UserPostHandler) Execute(ctx *gin.Context) {
	var userCreateParams core.UserCreateParams
	if err := ctx.ShouldBindJSON(&userCreateParams); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message := ctrl.useCase.Execute(userCreateParams)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
