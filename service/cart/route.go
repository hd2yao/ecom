package cart

import (
    "fmt"
    "net/http"

    "github.com/go-playground/validator/v10"
    "github.com/gorilla/mux"

    "github.com/hd2yao/ecom/service/auth"
    "github.com/hd2yao/ecom/types"
    "github.com/hd2yao/ecom/utils"
)

type Handler struct {
    store        types.OrderStore
    productStore types.ProductStore
    userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
    return &Handler{store: store, productStore: productStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handlerCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handlerCheckout(w http.ResponseWriter, r *http.Request) {
    userID := auth.GetUserIDFromContext(r.Context())

    var cart types.CartCheckoutPayload
    if err := utils.ParseJSON(r, &cart); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    if err := utils.Validate.Struct(cart); err != nil {
        errors := err.(validator.ValidationErrors)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
        return
    }

    // get products
    productIDs, err := getCartItemsIDs(cart.Items)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    products, err := h.productStore.GetProductsByIDs(productIDs)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
        "total_price": totalPrice,
        "order_id":    orderID,
    })
}
