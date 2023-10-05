package user

import (
	"context"
)

type UserService interface {
	List(ctx context.Context) ([]Provinces, error)
	Detail(ctx context.Context, id int64) (*Provinces, error)
	Create(ctx context.Context, input Provinces) (Provinces, error)
	Update(ctx context.Context, input Provinces) (*Provinces, error)
	Delete(ctx context.Context, id int64) error
}
