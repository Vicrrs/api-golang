package model

// Product e o modelo que representa um produto persistido e transportado pela API.
type Product struct {
	ID    int     `json:"id_products"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
