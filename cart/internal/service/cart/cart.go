package cart

import (
	"context"
	"errors"
	"fmt"
	"route256/cart/internal/model"
	"route256/cart/internal/repository"
	"route256/cart/internal/service/client/pim"
)

var ErrNotFound = errors.New("not found")
var ErrEmptyCart = errors.New("empty cart")
var ErrPIMProductNotFound = errors.New("pim product not found")
var ErrPIMLowAvailability = errors.New("pim product low availability in stock")

type PIMClient interface {
	GetProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error)
}

type LOMSClient interface {
	GetStockInfo(ctx context.Context, sku model.SKU) (uint64, error)
	CreateOrder(ctx context.Context, userID model.UserID, items []*model.Item) (model.OrderID, error)
}

type Cart struct {
	rep  repository.Cart
	pim  PIMClient
	loms LOMSClient
}

func New(rep repository.Cart, pim PIMClient, loms LOMSClient) *Cart {
	return &Cart{
		rep:  rep,
		pim:  pim,
		loms: loms,
	}
}

func (c *Cart) Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint16) error {
	_, err := c.pim.GetProductInfo(ctx, sku)
	if err != nil {
		if errors.Is(err, pim.ErrProductNotFound) {
			return ErrPIMProductNotFound
		}

		return err
	}

	available, err := c.loms.GetStockInfo(ctx, sku)
	if err != nil {
		return err
	}

	if uint64(count) > available {
		return ErrPIMLowAvailability
	}

	cartItem := model.Item{
		SKU:   sku,
		Count: count,
	}

	return c.rep.Add(ctx, userID, &cartItem)
}

func (c *Cart) Delete(ctx context.Context, userID model.UserID, sku model.SKU) error {
	err := c.rep.DeleteBySKU(ctx, userID, sku)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

// List returns a detailed list of the cart items.
// The items is a combination of cart values and product info from ProductService.
//
// TODO: the solution contains problem n+1 and incorrect behavior as a result of the context deadline.
func (c *Cart) List(ctx context.Context, userID model.UserID) ([]*model.ItemDetail, error) {
	list, err := c.rep.FindByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	detailList := make([]*model.ItemDetail, 0, len(list))
	for k := range list {
		productInfo, err := c.pim.GetProductInfo(ctx, list[k].SKU)
		if err != nil {
			return nil, fmt.Errorf("failed to get product info from product.service: %w", err)
		}

		detailList = append(detailList, &model.ItemDetail{
			SKU:   list[k].SKU,
			Count: list[k].Count,
			Price: productInfo.Price,
			Name:  productInfo.Name,
		})
	}

	return detailList, nil
}

func (c *Cart) Clear(ctx context.Context, userID model.UserID) error {
	err := c.rep.DeleteByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}

	return nil
}

func (c *Cart) Checkout(ctx context.Context, userID model.UserID) (model.OrderID, error) {
	items, err := c.rep.FindByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrEmptyCart
		}

		return 0, err
	}

	if len(items) == 0 {
		return 0, ErrEmptyCart
	}

	orderID, err := c.loms.CreateOrder(ctx, userID, items)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}
