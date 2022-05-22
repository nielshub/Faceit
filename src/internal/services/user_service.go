package service

import (
	"Faceit/src/internal/model"
	"Faceit/src/internal/ports"
	"context"
)

type UserService struct {
	relationalUserDBRepository ports.RelationalUserDBRepository
}

func NewUserService(relationalUserDBRepository ports.RelationalUserDBRepository) *UserService {
	return &UserService{
		relationalUserDBRepository: relationalUserDBRepository,
	}
}

func (user *UserService) CreateUser(ctx context.Context, newUser model.User) (string, error) {
	return "", nil
}

func (user *UserService) DeleteUser(ctx context.Context, userId string) error {
	return nil
}

func (user *UserService) UpdateUser(ctx context.Context, userId string, updatedUser model.User) error {
	return nil
}

func (user *UserService) GetUsers(ctx context.Context, filters map[string]string) ([]model.User, error) {
	return nil, nil
}
