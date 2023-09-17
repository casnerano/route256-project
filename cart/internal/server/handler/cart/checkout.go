package cart

import (
    "net/http"
    "time"
)

type checkoutHandler struct{}

func NewCheckoutHandler() *checkoutHandler {
    return &checkoutHandler{}
}

func (lh *checkoutHandler) Checkout(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    time.Sleep(time.Second)
}
