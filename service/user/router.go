package user

import (
    "fmt"
    "github.com/hd2yao/ecom/service/auth"
    "net/http"

    "github.com/gorilla/mux"

    "github.com/hd2yao/ecom/types"
    "github.com/hd2yao/ecom/utils"
)

type Handler struct {
    store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
    return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/login", h.handleLogin).Methods("POST")
    router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// 用户注册
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
    // get JSON payload
    var payload types.RegisterUserPayload
    if err := utils.ParseJSON(r, payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
    }

    // check if the user exists
    _, err := h.store.GetUserByEmail(payload.Email)
    if err == nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
        return
    }

    // if it doesn't, we create the new user
    hashedPassword, err := auth.HashPassword(payload.Password)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    err = h.store.CreateUser(types.User{
        FirstName: payload.FirstName,
        LastName:  payload.LastName,
        Email:     payload.Email,
        Password:  hashedPassword,
    })
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
    }
    utils.WriteJSON(w, http.StatusCreated, nil)
}
