package handler

import (
	"context"
	"encoding/json"
	"entdemo/ent"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

type IDB interface {
	GetAllUsers(ctx context.Context, name string) ([]*ent.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	Update(ctx context.Context, usr *ent.User) error
	GetUserFriends(ctx context.Context, id uuid.UUID) ([]*ent.User, error)
	UpdateFriends(ctx context.Context, id uuid.UUID, newFriendIds []uuid.UUID) error
}

type Handler struct {
	DB IDB
}

func NewHandler(db IDB) *Handler {
	return &Handler{db}
}

func (h *Handler) returnJson(w http.ResponseWriter, user interface{}) {
	if user == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		logrus.WithError(err).Error("error encoding user")
	}
}

func (h *Handler) returnOk(w http.ResponseWriter, user interface{}) {
	w.WriteHeader(http.StatusOK)
	h.returnJson(w, user)
}

func (h *Handler) returnCreated(w http.ResponseWriter, user interface{}) {
	w.WriteHeader(http.StatusCreated)
	h.returnJson(w, user)
}
