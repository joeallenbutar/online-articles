package impl

import (
	"context"
	"online-articles/common"
	"online-articles/entity"
	"online-articles/exception"
	"online-articles/model"
	"online-articles/repository"
	"online-articles/service"

	"golang.org/x/crypto/bcrypt"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) LoginAuthentication(ctx context.Context, userModel model.UserModel) entity.User {
	common.Validate(userModel)
	userResult, err := userService.UserRepository.LoginAuthentication(ctx, userModel.Username)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}
	err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(userModel.Password))
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: "incorrect username and password",
		})
	}
	common.NewLogger().Info("UserService-LoginAuthentication - username: ", userModel.Username)

	return userResult
}

func (service *userServiceImpl) RegisterUser(ctx context.Context, userModel model.UserModel) model.UserModel {
	common.Validate(userModel)
	hashPassword, _ := HashPassword(userModel.Password)
	user := entity.User{
		Username: userModel.Username,
		Password: hashPassword,
	}

	result := service.UserRepository.CreateUser(ctx, user)
	common.NewLogger().Info("UserService-RegisterUser - username: ", userModel.Username)
	userModel = model.UserModel{
		Username: result.Username,
		Password: userModel.Password,
	}
	return userModel
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}
