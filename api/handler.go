package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/luquxSentinel/spacedrive/service"
	"github.com/luquxSentinel/spacedrive/types"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: authService,
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	// handle sign up

	reqData := new(types.CreateUserData)

	err := decodeJSON(r.Body, reqData)
	if err != nil {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}

	// call auth service create user
	err = h.service.CreateUser(r.Context(), reqData)
	if err != nil {
		// TODO: implement logger service
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// handle sign in
	reqData := new(types.LoginData)

	err := decodeJSON(r.Body, reqData)
	if err != nil {
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}

	// call auth service signin service

	user, token, err := h.service.SignIn(r.Body, reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write response status
	w.WriteHeader(http.StatusOK)

	// set authorization header
	w.Header().Set("authorization", token)

	// write response
	err =
		writeJSON(w, user)
	if err != nil {
		http.Error(w, "response error. please try again", http.StatusInternalServerError)
		return
	}

}

func decodeJSON(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func writeJSON(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}
