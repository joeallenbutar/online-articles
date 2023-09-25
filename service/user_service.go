package service

import (
	"context"
	"online-articles/entity"
	"online-articles/model"
)

type UserService interface {
	LoginAuthentication(ctx context.Context, model model.UserModel) entity.User
	RegisterUser(ctx context.Context, model model.UserModel) model.UserModel
}
