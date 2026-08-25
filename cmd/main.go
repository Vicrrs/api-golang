package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

// main inicializa o servidor HTTP e registra suas rotas.
func main() {

	server := gin.Default()

	// Camada de controllers
	ProductController := controller.NewProductController()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Mapeando as rotas
	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
