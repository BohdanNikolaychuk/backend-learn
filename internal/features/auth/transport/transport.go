package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/BohdanNikolaychuk/backend-learn/internal/core/tools"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/auth/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
	Refresh(ctx context.Context, req dto.RefreshRequest) (*dto.LoginResponse, error)
	Logout(ctx context.Context, req dto.RefreshRequest) error
}

type AuthTransport struct {
	service AuthService
}

func NewAuthTransport(service AuthService) *AuthTransport {
	return &AuthTransport{service: service}
}

func (t *AuthTransport) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	tokens, err := t.service.Login(r.Context(), req)

	if err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := dto.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	tools.WriteJSON(w, http.StatusOK, resp)
}

func (t *AuthTransport) Refresh(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	tokens, err := t.service.Refresh(r.Context(), req)
	if err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	resp := dto.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	tools.WriteJSON(w, http.StatusOK, resp)
}

func (t *AuthTransport) Logout(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := t.service.Logout(r.Context(), req)
	if err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
