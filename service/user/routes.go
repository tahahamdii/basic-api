package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tahahamdii/basic-api/service/auth"
	"github.com/tahahamdii/basic-api/types"
	"github.com/tahahamdii/basic-api/utils"
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
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	//get json payload
	
	var payload  types.RegisterUserPayload
	if err := utils.ParseJson(r,payload); err != nil {
		utils.WriteError(w,http.StatusBadRequest, err)
	}
	
	//check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists",payload.Email))
		return
	}


	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}
	// if user doesnt exists

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
