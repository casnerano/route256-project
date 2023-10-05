package model

type Stock struct {
	ID        int
	SKU       SKU
	Available uint32
	Reserved  uint32
}
