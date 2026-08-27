package usecase

import (
	"go-api/model"
	"go-api/repository"
)

// ProductUsecase coordena as regras de negocio entre os handlers e o repositorio de produtos.
type ProductUsecase struct {
	repository repository.ProductRepository
}

// NewProductUseCase cria um caso de uso de produtos vinculado ao repositorio informado.
func NewProductUseCase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

// GetProducts solicita ao repositorio a lista de produtos disponiveis.
func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

// CreateProduct valida o fluxo de criacao, persiste o produto e associa o identificador gerado.
func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil

}

// GetProdutctById solicita ao repositorio o produto correspondente ao identificador informado.
func (pu *ProductUsecase) GetProdutctById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProdutctById(id_product)

	if err != nil {
		return nil, err
	}
	return product, nil
}
