package cart

import (
    "net/http"

    "github.com/gorilla/mux"

    "github.com/hd2yao/ecom/types"
)

type Handler struct {
    store types.OrderStore
}

func NewHandler(store types.OrderStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/cart/checkout", h.handlerCheckout).Methods(http.MethodPost)
}

func (h *Handler) handlerCheckout(w http.ResponseWriter, r *http.Request) {

}
