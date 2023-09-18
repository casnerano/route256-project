package product

type GetProductRequest struct {
    Token string `json:"token"`
    SKU   uint32 `json:"sku"`
}

type GetProductResponse struct {
    Name  string `json:"name"`
    Price uint32 `json:"price"`
}

type GetProductErrorResponse struct {
    Code    int               `json:"code"`
    Message string            `json:"message"`
    Details map[string]string `json:"details,omitempty"`
}
