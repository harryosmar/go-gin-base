package user

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/go_gin_base"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&UserModel{})

	DB = database
}
