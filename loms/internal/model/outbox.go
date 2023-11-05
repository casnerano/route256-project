package model

import "time"

type OrderStatusOutbox struct {
	ID          int
	OrderID     OrderID
	OrderStatus OrderStatus
	IsDelivery  bool
	CreatedAt   time.Time
}
