package auth

import (
	"AustemServer/models"
	"context"
)

type Repository interface {
	Get(ctx context.Context, username, password string) (*models.Users, error)
}
