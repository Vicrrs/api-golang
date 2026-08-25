package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// productController centraliza o tratamento das requisições de produtos.
type productController struct {
	// use case
}

// NewProductController cria e retorna um controller de produtos.
func NewProductController() productController {
	return productController{}
}

// GetProducts responde com a lista de produtos disponíveis.
func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata Frita",
			Price: 20,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
