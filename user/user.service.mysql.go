package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserServiceMySQL struct {
	db *gorm.DB
}

func (u *UserServiceMySQL) Index(ctx context.Context) ([]Provinces, error) {
	var provinces []Provinces //

	if err := u.db.Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}

func NewUserServiceMySQL(database *gorm.DB) *UserServiceMySQL {
	return &UserServiceMySQL{
		db: database,
	}
}

func (u *UserServiceMySQL) List(ctx context.Context) ([]Provinces, error) {
	var provinces []Provinces //

	if err := u.db.Find(&provinces).Error; err != nil {
		return nil, err
	}

	return provinces, nil
}

func (u *UserServiceMySQL) Detail(ctx context.Context, id int64) (*Provinces, error) {
	var provinces Provinces

	if err := u.db.Where("id = ?", id).First(&provinces).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, err
		}
		return nil, err
	}

	return &provinces, nil
}

func (u *UserServiceMySQL) Create(ctx context.Context, input Provinces) (Provinces, error) {
	var provinces Provinces

	if err := u.db.Create(&input).Error; err != nil {

		return provinces, err
	}

	return input, nil
}

func (u *UserServiceMySQL) Update(ctx context.Context, input Provinces) (*Provinces, error) {
	if err := u.db.Model(&Provinces{}).Where("id = ?", input.Id).Updates(input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (u *UserServiceMySQL) Delete(ctx context.Context, id int64) error {
	if err := u.db.Where("id = ?", id).Delete(&Provinces{}).Error; err != nil {
		return err
	}
	return nil
}
