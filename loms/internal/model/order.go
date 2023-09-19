package model

import "time"

type OrderID = int64

type OrderStatus string

const (
	OrderStatusNew          OrderStatus = "NEW"
	OrderStatusAwaitPayment OrderStatus = "AWAIT PAYMENT"
	OrderStatusFailed       OrderStatus = "FAILED"
	OrderStatusPayed        OrderStatus = "PAYED"
	OrderStatusCanceled     OrderStatus = "CANCELED"
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
