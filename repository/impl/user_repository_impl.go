package impl

import (
	"context"
	"errors"
	"online-articles/entity"
	"online-articles/exception"
	"online-articles/repository"

	"gorm.io/gorm"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) CreateUser(ctx context.Context, user entity.User) entity.User {
	user = entity.User{
		Username: user.Username,
		Password: user.Password,
		IsActive: 1,
	}
	err := userRepository.DB.WithContext(ctx).Create(&user).Error
	exception.PanicLogging(err)
	return user
}

func (userRepository *userRepositoryImpl) LoginAuthentication(ctx context.Context, username string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).Where("username = ? and is_active = ?", username, 1).Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
