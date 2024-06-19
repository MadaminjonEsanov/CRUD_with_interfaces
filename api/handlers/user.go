package handlers

import (
	"app/models"
	"context"
)

var ctx = context.Background()

func (h *Handlers) CreateUser() {
	h.storage.GetUserRepo().CreateUser(ctx, models.User{})
}