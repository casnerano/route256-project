package order

import (
	"route256/loms/internal/model"
)

type item struct {
	SKU   model.SKU `json:"sku"`
	Count uint64    `json:"count"`
}

type createRequest struct {
	User  model.UserID `json:"user"`
	Items []item       `json:"items"`
}

type createResponse struct {
	OrderID model.OrderID `json:"orderID"`
}
