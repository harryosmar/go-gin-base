package user

import "context"

type UserServiceMySQL struct {
}

func NewUserServiceMySQL() *UserServiceMySQL {
	return &UserServiceMySQL{}
}

func (u UserServiceMySQL) List(ctx context.Context, page int, limit int) ([]UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMySQL) Detail(ctx context.Context, id int64) (*UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMySQL) Create(ctx context.Context, input UserModel) (UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMySQL) Update(ctx context.Context, input UserModel) (*UserModel, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMySQL) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
