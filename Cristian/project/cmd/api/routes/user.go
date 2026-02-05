package routes

import (
	"api-gateway/cmd/api/handlers"
	application "api-gateway/pkg/application/user"
	infrastructure "api-gateway/pkg/infrastructure/postgresql"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {

	userRepository := infrastructure.NewPostgresRepository("users", db)
	useCase := application.NewUserCreator(userRepository)
	handler := handlers.NewUserPostHandler(*useCase)
	router.POST("/users", handler.Execute)
}
