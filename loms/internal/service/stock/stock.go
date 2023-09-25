package stock

import (
	"context"
	"route256/loms/internal/model"
	"route256/loms/internal/repository"
)

type Stock struct {
	rep repository.Stock
}

func New(rep repository.Stock) *Stock {
	return &Stock{rep: rep}
}

func (s *Stock) GetAvailable(ctx context.Context, sku model.SKU) (uint32, error) {
	st, err := s.rep.FindBySKU(ctx, sku)
	if err != nil {
		return 0, err
	}
	return st.Available, nil
}

func (s *Stock) AddReserve(ctx context.Context, sku model.SKU, count uint32) error {
	return s.rep.AddReserve(ctx, sku, count)
}

func (s *Stock) CancelReserve(ctx context.Context, sku model.SKU, count uint32) error {
	return s.rep.CancelReserve(ctx, sku, count)
}

func (s *Stock) ShipReserve(ctx context.Context, sku model.SKU, count uint32) error {
	return s.rep.ShipReserve(ctx, sku, count)
}
