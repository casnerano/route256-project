package loms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"

	"route256/cart/internal/model"
)

type client struct {
	baseURL string
}

func NewClient(baseURL string) *client {
	return &client{baseURL: baseURL}
}

func (c *client) GetStockInfo(ctx context.Context, sku model.SKU) (uint64, error) {
	path, err := url.JoinPath(c.baseURL, "/api/stock/info")
	if err != nil {
		return 0, err
	}

	requestPayload := GetStockInfoRequest{SKU: sku}

	bRequestPayload, err := json.Marshal(requestPayload)

	if err != nil {
		return 0, err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		path,
		bytes.NewReader(bRequestPayload),
	)
	if err != nil {
		return 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, fmt.Errorf("failed request product stock info: %w", err)
	}

	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Printf("Failed close response body: %s\n", debug.Stack())
		}
	}()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed request, server returned response code %d", response.StatusCode)
	}

	responsePayload := GetStockInfoResponse{}
	if err = json.NewDecoder(response.Body).Decode(&responsePayload); err != nil {
		return 0, err
	}

	return responsePayload.Count, nil
}
