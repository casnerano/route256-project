package model

import (
	"fmt"
	"time"
)

type OrderID = uint64

type OrderStatus string

const (
	OrderStatusNew          OrderStatus = "new"
	OrderStatusAwaitPayment OrderStatus = "await_payment"
	OrderStatusFailed       OrderStatus = "failed"
	OrderStatusPayed        OrderStatus = "payed"
	OrderStatusCanceled     OrderStatus = "canceled"
)

func (status *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
	case string:
		*status = OrderStatus(s)
	default:
		return fmt.Errorf("unknown order status: %T", src)
	}
	return nil
}

type OrderItem struct {
	SKU   SKU
	Count uint32
}

type Order struct {
	ID        OrderID
	Status    OrderStatus
	User      UserID
	Items     []*OrderItem
	CreatedAt time.Time `json:"created_at"`
}
