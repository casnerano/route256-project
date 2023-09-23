package pim

import "route256/cart/internal/model"

type GetProductRequest struct {
	Token string    `json:"token"`
	SKU   model.SKU `json:"sku"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type GetProductErrorResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}
