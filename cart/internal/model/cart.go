package model

type Item struct {
	SKU   SKU
	Count uint16
}

type ItemDetail struct {
	SKU   SKU
	Count uint16
	Price uint32
	Name  string
}

type Cart struct {
	UserID UserID
	Items  []Item
}
