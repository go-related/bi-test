package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserDTO struct {
	ID *uuid.UUID `json:"id,omitempty"`

	Age       int       `json:"age,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	GroupID *uuid.UUID `json:"group_id,omitempty"`
	Friends []*UserDTO `json:"friends,omitempty"`
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	nameParameter := r.URL.Query().Get("name")
	users, err := h.DB.GetAllUsers(r.Context(), nameParameter)
	if err != nil {
		logrus.WithError(err).Error("error getting users")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	usersJson, err := json.Marshal(users)
	if err != nil {
		logrus.WithError(err).Error("error marshalling users")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.returnOk(w, usersJson)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := GetParameter(r, "userID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}
	user, err := h.DB.GetUserById(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	h.returnOk(w, user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userModel := convertUserToDomainModel(&user)
	userCreated, err := h.DB.Create(r.Context(), userModel)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}
	h.returnCreated(w, userCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := GetParameter(r, "userID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}
	var user UserDTO
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	user.ID = &id
	userModel := convertUserToDomainModel(&user)
	if err := h.DB.Update(r.Context(), userModel); err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}
	h.returnOk(w, user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := GetParameter(r, "userID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}
	if err := h.DB.Delete(r.Context(), id); err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetUserFriends(w http.ResponseWriter, r *http.Request) {
	idStr := GetParameter(r, "userID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}
	users, err := h.DB.GetUserFriends(r.Context(), id)
	if err != nil {
		http.Error(w, "users friends could not be found", http.StatusInternalServerError)
		return
	}
	h.returnOk(w, users)
}

func (h *Handler) UpdateUserFriends(w http.ResponseWriter, r *http.Request) {
	idStr := GetParameter(r, "userID")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	var newFriends []uuid.UUID
	err = json.NewDecoder(r.Body).Decode(&newFriends)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.DB.UpdateFriends(r.Context(), id, newFriends)
	if err != nil {
		http.Error(w, "failed to update users friends", http.StatusInternalServerError)
		return
	}
	h.returnOk(w, nil)
}
