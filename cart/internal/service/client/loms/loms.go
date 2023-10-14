package loms

import (
	"context"
	"route256/cart/internal/model"
	pbOrder "route256/loms/pkg/proto/order/v1"
	pbStock "route256/loms/pkg/proto/stock/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcConn    *grpc.ClientConn
	stockClient pbStock.StockClient
	orderClient pbOrder.OrderClient
}

func NewClient(addr string) (*Client, error) {
	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcConn:    grpcConn,
		stockClient: pbStock.NewStockClient(grpcConn),
		orderClient: pbOrder.NewOrderClient(grpcConn),
	}, nil
}

func (c *Client) GetStockInfo(ctx context.Context, sku model.SKU) (uint32, error) {
	infoResponse, err := c.stockClient.Info(ctx, &pbStock.InfoRequest{Sku: sku})
	if err != nil {
		return 0, err
	}

	return infoResponse.GetCount(), nil
}

func (c *Client) CreateOrder(ctx context.Context, userID model.UserID, items []*model.Item) (model.OrderID, error) {
	pbItems := make([]*pbOrder.Item, 0)
	for _, item := range items {
		pbItems = append(pbItems, &pbOrder.Item{
			Sku:   item.SKU,
			Count: item.Count,
		})
	}

	createResponse, err := c.orderClient.Create(ctx, &pbOrder.CreateRequest{
		User:  userID,
		Items: pbItems,
	})
	if err != nil {
		return 0, err
	}

	return createResponse.GetOrderId(), nil
}

func (c *Client) Close() error {
	return c.grpcConn.Close()
}
