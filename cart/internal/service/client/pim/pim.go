package pim

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"route256/cart/internal/model"
	"runtime/debug"
)

var ErrProductNotFound = errors.New("product not found")

type Client struct {
	baseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{baseURL: baseURL}
}

// TODO: token from context
func (c *Client) GetProductInfo(ctx context.Context, sku model.SKU) (*model.ProductInfo, error) {
	path, err := url.JoinPath(c.baseURL, "/get_product")
	if err != nil {
		return nil, err
	}

	requestPayload := GetProductRequest{
		Token: "testtoken",
		SKU:   sku,
	}

	bRequestPayload, err := json.Marshal(requestPayload)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		path,
		bytes.NewReader(bRequestPayload),
	)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed request get product info: %w", err)
	}

	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Printf("Failed close response body: %s\n", debug.Stack())
		}
	}()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, ErrProductNotFound
		}

		errResponse := GetProductErrorResponse{}
		if err = json.NewDecoder(response.Body).Decode(&errResponse); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("failed request, server returned response code %d", response.StatusCode)
	}

	responsePayload := GetProductResponse{}
	if err = json.NewDecoder(response.Body).Decode(&responsePayload); err != nil {
		return nil, err
	}

	productInfo := &model.ProductInfo{
		Name:  responsePayload.Name,
		Price: responsePayload.Price,
	}

	return productInfo, nil
}
