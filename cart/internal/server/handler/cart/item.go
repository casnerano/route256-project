package cart

import (
    "net/http"
    "time"
)

type requestItem struct {
    User uint64
    SKU  uint64
}

type itemHandler struct{}

func NewItemHandler() *itemHandler {
    return &itemHandler{}
}

func (ih *itemHandler) Add(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    time.Sleep(time.Second)
}

func (ih *itemHandler) Delete(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    time.Sleep(time.Second)
}
