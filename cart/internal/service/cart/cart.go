package cart

import (
	"context"
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
)

type PIMClient interface {
	GetProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error)
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

func (c *cart) Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint16) error {
	_, err := c.pim.GetProductInfo(ctx, sku)
	if err != nil {
		return err
	}

	cartItem := model.CartItem{
		SKU:   sku,
		Count: count,
	}

	// todo: check stocks

	return c.rep.Add(ctx, userID, &cartItem)
}

func (c *cart) Delete(ctx context.Context, userID model.UserID, sku model.SKU) error {
	return c.rep.DeleteBySKU(ctx, userID, sku)
}

// List returns a detailed list of the cart items.
// The items is a combination of cart values and product info from ProductService.
//
// TODO: the solution contains problem n+1 and incorrect behavior as a result of the context deadline.
func (c *cart) List(ctx context.Context, userID model.UserID) ([]*model.CartItemDetail, error) {
	list, err := c.rep.FindByUser(ctx, userID)
	if err != nil && err != repository.ErrNotFound {
		return nil, err
	}

	detailList := make([]*model.CartItemDetail, 0, len(list))
	for k := range list {
		productInfo, err := c.pim.GetProductInfo(ctx, list[k].SKU)
		if err != nil {
			return nil, fmt.Errorf("failed to get product info from product.service: %w", err)
		}

		detailList = append(detailList, &model.CartItemDetail{
			SKU:   list[k].SKU,
			Count: list[k].Count,
			Price: productInfo.Price,
			Name:  productInfo.Name,
		})
	}

	return detailList, nil
}

func (c *cart) Clear(ctx context.Context, userID model.UserID) error {
	return c.rep.DeleteByUser(ctx, userID)
}
