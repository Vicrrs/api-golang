package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

// main e o ponto de entrada que conecta o banco, monta as camadas da aplicacao e inicia o servidor HTTP.
func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Mapeando as rotas
	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProdutctById)
	server.Run(":8000")
}
