package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserServiceMySQL struct {
	db *gorm.DB
}

func NewUserServiceMySQL(database *gorm.DB) *UserServiceMySQL {
	return &UserServiceMySQL{
		db: database, // Inisialisasi db menggunakan parameter yang diterima saat pembuatan instance.
	}
}

func (u *UserServiceMySQL) List(ctx context.Context, page int, limit int) ([]UserModel, error) {
	var usermodels []UserModel

	if err := u.db.Find(&usermodels).Error; err != nil { // Menggunakan u.db
		return nil, err
	}

	return usermodels, nil
}

func (u *UserServiceMySQL) Detail(ctx context.Context, id int64) (*UserModel, error) {
	var userModel UserModel

	if err := u.db.Where("id = ?", id).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, err
		}
		return nil, err
	}

	return &userModel, nil
}

func (u *UserServiceMySQL) Create(ctx context.Context, input UserModel) (UserModel, error) {
	var userModel UserModel

	if err := u.db.Create(&input).Error; err != nil {

		return userModel, err
	}

	return input, nil
}

func (u *UserServiceMySQL) Update(ctx context.Context, input UserModel) (*UserModel, error) {
	if err := u.db.Model(&UserModel{}).Where("id = ?", input.Id).Updates(input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (u *UserServiceMySQL) Delete(ctx context.Context, id int64) error {
	if err := u.db.Where("id = ?", id).Delete(&UserModel{}).Error; err != nil {
		return err
	}
	return nil
}
