package cart

import (
	"net/http"
	"time"
)

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	time.Sleep(time.Second)
}
