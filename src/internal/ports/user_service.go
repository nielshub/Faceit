package ports

import (
	"Faceit/src/internal/model"
	"context"
)

type UserService interface {
	CreateUser(context.Context, model.User) (string, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, model.User) error
	GetUsers(context.Context, map[string]string) ([]model.User, error)
}
