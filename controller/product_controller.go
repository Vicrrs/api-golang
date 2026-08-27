package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// productController e o controlador HTTP que recebe e responde as requisicoes relacionadas a produtos.
type productController struct {
	productUseCase usecase.ProductUsecase
}

// NewProductController e o construtor que cria um controlador de produtos com seu caso de uso.
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// GetProducts e o handler que busca e retorna todos os produtos em uma resposta HTTP.
func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// CreateProduct e o handler que recebe os dados de um produto, salva-o e retorna o produto criado.
func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

// GetProdutctById e o handler que busca um produto pelo identificador informado na URL.
func (p *productController) GetProdutctById(ctx *gin.Context) {

	id := ctx.Param("productId")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUseCase.GetProdutctById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado na base de dados.",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// DeletProduct remove o produto cujo ID foi informado na URL e retorna o resultado da operação
func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto não pode ser nulo.",
		})
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Id do produto precisa ser um número.",
		})
		return
	}

	deleted, err := p.productUseCase.DeleteProduct(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if !deleted {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "Produto não foi encontrado na base de dados.",
		})
		return
	}
	ctx.Status(http.StatusNoContent)
}
