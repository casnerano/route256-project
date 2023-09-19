package loms

import "route256/cart/internal/model"

type GetStockInfoRequest struct {
	SKU model.SKU `json:"sku"`
}

type GetStockInfoResponse struct {
	Count uint64 `json:"count"`
}
