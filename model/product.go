package model

// Product representa um produto e seus dados expostos pela API.
type Product struct {
	ID    int     `json:"id_products"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
