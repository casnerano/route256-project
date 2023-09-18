package model

type CartItem struct {
	SKU   uint32
	Name  string
	Price uint32
	Count uint16
}

type Cart struct {
	UserID UserID
	Items  []CartItem
}
