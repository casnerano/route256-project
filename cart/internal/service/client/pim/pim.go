package pim

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256/cart/internal/model"
	"route256/cart/pkg/interceptor"
	pb "route256/cart/pkg/proto/client/product"
)

var ErrProductNotFound = errors.New("product not found")

type Client struct {
	grpcConn      *grpc.ClientConn
	productClient pb.ProductServiceClient
}

func NewClient(addr string) (*Client, error) {
	grpcConn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.ClientUnaryRateLimiter()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcConn:      grpcConn,
		productClient: pb.NewProductServiceClient(grpcConn),
	}, nil
}

// TODO: token from context
func (c *Client) GetProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error) {
	requestPayload := pb.GetProductRequest{
		Token: "testtoken",
		Sku:   uint32(sku),
	}

	response, err := c.productClient.GetProduct(ctx, &requestPayload)
	if err != nil {
		return nil, fmt.Errorf("failed request get product info: %w", err)
	}

	productInfo := &model.ProductInfo{
		Name:  response.GetName(),
		Price: response.GetPrice(),
	}

	return productInfo, nil
}

func (c *Client) Close() error {
	return c.grpcConn.Close()
}
