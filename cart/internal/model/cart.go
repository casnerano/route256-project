package model

type Item struct {
	ID     int
	UserID UserID
	SKU    SKU
	Count  uint32
}

type ItemDetail struct {
	Item
	Price uint32
	Name  string
}

type Cart struct {
	UserID UserID
	Items  []Item
}
