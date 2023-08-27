package user

import "time"

type UserModel struct {
	Id          int64     `json:"id" uri:"id"`
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type Paging struct {
	Page  int `uri:"page" `
	Limit int `uri:"limit"`
}
