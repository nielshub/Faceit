package ports

import (
	"Faceit/src/internal/model"
	"context"
)

type RelationalUserDBRepository interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, string, model.User) (*model.User, error)
	GetUsers(context.Context, string, string) ([]model.User, error)
}
