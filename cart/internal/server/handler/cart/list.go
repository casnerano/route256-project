package cart

import (
    "net/http"
    "time"
)

type listHandler struct{}

func NewListHandler() *listHandler {
    return &listHandler{}
}

func (lh *listHandler) List(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    time.Sleep(time.Second)
}
