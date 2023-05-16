package auth

import (
	"AustemServer/models"
	"context"
)

type UseCase interface {
	SignIn(ctx context.Context, user *models.Users) (string, error)
}
