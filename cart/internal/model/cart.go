package model

type SKU = uint64

type CartItem struct {
	SKU   SKU
	Count uint16
}

type CartItemDetail struct {
	SKU   SKU
	Count uint16
	Price uint32
	Name  string
}

type Cart struct {
	UserID UserID
	Items  []CartItem
}
