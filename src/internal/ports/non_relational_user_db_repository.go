package ports

import (
	"Faceit/src/internal/model"
	"context"
)

type NonRelationalUserDBRepository interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, *model.User) (*model.User, error)
	GetUsersWithFilter(context.Context, string, string) ([]model.User, error)
	GetAllUsers(context.Context) ([]model.User, error)
}
