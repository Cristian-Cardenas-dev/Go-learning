package main

import (
	"api-gateway/cmd/api/routes"
	infrastructure "api-gateway/pkg/infrastructure/postgresql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := infrastructure.NewPostgresConnection()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		db.Close()
		panic("failed")
	}
	routes.RegisterRoutes(r, db)
	r.Run(":8080")
}
