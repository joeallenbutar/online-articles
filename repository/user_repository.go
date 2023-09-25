package repository

import (
	"context"
	"online-articles/entity"
)

type UserRepository interface {
	LoginAuthentication(ctx context.Context, username string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) entity.User
}
