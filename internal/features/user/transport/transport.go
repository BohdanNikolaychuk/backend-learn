package transport

import (
	"encoding/json"
	"net/http"

	"github.com/BohdanNikolaychuk/backend-learn/internal/core/tools"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/dto"
	"github.com/BohdanNikolaychuk/backend-learn/internal/features/user/service"
)

type UserTransport struct {
	service *service.UserService
}

func NewUserTransport(service *service.UserService) *UserTransport {
	return &UserTransport{service: service}
}

func (t *UserTransport) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := t.service.Create(r.Context(), req)
	if err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp := dto.CreateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	tools.WriteJSON(w, http.StatusOK, resp)
}

func (t *UserTransport) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		tools.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := t.service.GetByID(r.Context(), id)
	if err != nil {
		tools.WriteError(w, http.StatusNotFound, "user not found")
		return
	}

	resp := dto.CreateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	tools.WriteJSON(w, http.StatusOK, resp)
}

func (t *UserTransport) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		tools.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	_, err := t.service.DeleteByID(r.Context(), id)
	if err != nil {
		tools.WriteError(w, http.StatusNotFound, "user not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (t *UserTransport) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		tools.WriteError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		tools.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := t.service.Update(r.Context(), id, req)
	if err != nil {
		tools.WriteError(w, http.StatusNotFound, "user not found")
		return
	}

	resp := dto.CreateUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FullName:  user.FullName,
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	tools.WriteJSON(w, http.StatusOK, resp)
}
