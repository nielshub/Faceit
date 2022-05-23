package service

import (
	"Faceit/src/internal/model"
	"Faceit/src/internal/ports"
	"context"
)

type UserService struct {
	relationalUserDBRepository ports.NonRelationalUserDBRepository
}

func NewUserService(relationalUserDBRepository ports.NonRelationalUserDBRepository) *UserService {
	return &UserService{
		relationalUserDBRepository: relationalUserDBRepository,
	}
}

func (user *UserService) CreateUser(ctx context.Context, newUser model.User) (*model.User, error) {
	return user.relationalUserDBRepository.CreateUser(ctx, &newUser)
}

func (user *UserService) DeleteUser(ctx context.Context, userId string) error {
	return user.relationalUserDBRepository.DeleteUser(ctx, userId)
}

func (user *UserService) UpdateUser(ctx context.Context, userId string, updatedUser model.User) (*model.User, error) {
	return user.relationalUserDBRepository.UpdateUser(ctx, userId, &updatedUser)
}

func (user *UserService) GetUsers(ctx context.Context, key, value string) ([]model.User, error) {
	return user.relationalUserDBRepository.GetUsers(ctx, key, value)
}
