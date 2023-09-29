package db

import (
	"log"

	"github.com/harryosmar/go-gin-base/user"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}

	dbHost := viper.GetString("db_host")

	dbPort := viper.GetString("db_port")

	dbUser := viper.GetString("db_user")

	dbPassword := viper.GetString("db_password")

	dbName := viper.GetString("db_name")

	dbConnectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	database, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	database.AutoMigrate(&user.UserModel{})

	return database

}
