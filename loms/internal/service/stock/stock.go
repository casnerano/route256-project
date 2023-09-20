package stock

import (
	"context"

	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

type stock struct {
	rep repository.Stock
}

func New(rep repository.Stock) *stock {
	return &stock{rep: rep}
}

func (s *stock) GetAvailable(ctx context.Context, sku model.SKU) (uint16, error) {
	st, err := s.rep.FindBySKU(ctx, sku)
	if err != nil {
		return 0, err
	}
	return st.Available, nil
}

func (s *stock) AddReserve(ctx context.Context, sku model.SKU, count uint16) error {
	return s.rep.AddReserve(ctx, sku, count)
}

func (s *stock) CancelReserve(ctx context.Context, sku model.SKU, count uint16) error {
	return s.rep.CancelReserve(ctx, sku, count)
}

func (s *stock) ShipReserve(ctx context.Context, sku model.SKU, count uint16) error {
	return s.rep.ShipReserve(ctx, sku, count)
}
