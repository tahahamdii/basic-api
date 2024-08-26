package user

import (
	"encoding/json"
	"net/http"
	"github.com/tahahamdii/basic-api/types"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

	//get json payload
	if r.Body == nil {

	}
	var payload  types.RegisterUserPayload
	err := json.NewDecoder(r.Body).Decode(payload)
}
