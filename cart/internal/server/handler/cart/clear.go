package cart

import (
    "net/http"
    "time"
)

type clearHandler struct{}

func NewClearHandler() *clearHandler {
    return &clearHandler{}
}

func (lh *clearHandler) Clear(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    time.Sleep(time.Second)
}
