package ports

import (
	"Faceit/src/internal/model"
	"context"
)

type RelationalUserDBRepository interface {
	Create(context.Context, model.User) (string, error)
	Delete(context.Context, string) error
	Update(context.Context, string, model.User) error
	Get(context.Context, map[string]string) ([]model.User, error)
}
