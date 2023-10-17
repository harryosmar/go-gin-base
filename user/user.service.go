package user

import (
	"context"
)

type UserService interface {
	IndexProvinces(ctx context.Context) ([]Provinces, error)
	IndexRegencies(ctx context.Context) ([]Regencies, error)
	IndexDistricts(ctx context.Context) ([]Districts, error)
	IndexVillages(ctx context.Context) ([]Villages, error)
	List(ctx context.Context) ([]Provinces, error)
	Detail(ctx context.Context, id int64) (*Provinces, error)
	Create(ctx context.Context, input Provinces) (Provinces, error)
	Update(ctx context.Context, input Provinces) (*Provinces, error)
	Delete(ctx context.Context, id int64) error
}
