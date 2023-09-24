package loms

import "route256/cart/internal/model"

type GetStockInfoRequest struct {
	SKU model.SKU `json:"sku"`
}

type GetStockInfoResponse struct {
	Count uint64 `json:"count"`
}

type CreateOrderItem struct {
	SKU   model.SKU `json:"sku"`
	Count uint32    `json:"count"`
}

type CreateOrderRequest struct {
	User  model.UserID       `json:"user"`
	Items []*CreateOrderItem `json:"items"`
}

type CreateOrderResponse struct {
	OrderID model.OrderID `json:"orderID"`
}
