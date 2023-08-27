package user

import (
	"context"
)

type UserService interface {
	List(ctx context.Context, page int, limit int) ([]UserModel, error)
	Detail(ctx context.Context, id int64) (*UserModel, error)
	Create(ctx context.Context, input UserModel) (UserModel, error)
	Update(ctx context.Context, input UserModel) (*UserModel, error)
	Delete(ctx context.Context, id int64) error
}
