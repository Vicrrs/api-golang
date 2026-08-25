package controller

import (
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// productController centraliza o tratamento das requisições de produtos.
type productController struct {
	productUseCase usecase.ProductUsecase
}

// NewProductController cria e retorna um controller de produtos.
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// GetProducts responde com a lista de produtos disponíveis.
func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
