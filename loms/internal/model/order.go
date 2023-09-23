package model

import "time"

type OrderID = int64

type OrderStatus string

const (
	OrderStatusNew          OrderStatus = "new"
	OrderStatusAwaitPayment OrderStatus = "await payment"
	OrderStatusFailed       OrderStatus = "failed"
	OrderStatusPayed        OrderStatus = "payed"
	OrderStatusCanceled     OrderStatus = "canceled"
)

type OrderItem struct {
	SKU   SKU
	Count uint16
}

type Order struct {
	ID        OrderID
	Status    OrderStatus
	User      UserID
	Items     []*OrderItem
	CreatedAt time.Time `json:"created_at"`
}
