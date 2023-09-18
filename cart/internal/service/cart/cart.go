package cart

import (
	"context"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
)

type PIMClient interface {
	GetProductInfo(ctx context.Context, sku uint32) (*model.ProductInfo, error)
}

type LOMSClient interface{}

type cart struct {
	rep  repository.Cart
	pim  PIMClient
	loms LOMSClient
}

func New(rep repository.Cart, pim PIMClient, loms LOMSClient) *cart {
	return &cart{
		rep:  rep,
		pim:  pim,
		loms: loms,
	}
}

func (c *cart) Add(ctx context.Context, userID model.UserID, sku uint32, count uint16) error {
	productInfo, err := c.pim.GetProductInfo(ctx, sku)
	if err != nil {
		return err
	}

	cartItem := model.CartItem{
		SKU:   sku,
		Name:  productInfo.Name,
		Price: productInfo.Price,
		Count: count,
	}

	// todo: check stocks

	return c.rep.Add(ctx, userID, &cartItem)
}

func (c *cart) Delete(ctx context.Context, userID model.UserID, sku uint32) error {
	return c.rep.DeleteBySKU(ctx, userID, sku)
}

func (c *cart) List(ctx context.Context, userID model.UserID) ([]*model.CartItem, error) {
	list, err := c.rep.FindByUser(ctx, userID)
	if err != nil && err != repository.ErrNotFound {
		return nil, err
	}
	return list, nil
}

func (c *cart) Clear(ctx context.Context, userID model.UserID) error {
	return c.rep.DeleteByUser(ctx, userID)
}
