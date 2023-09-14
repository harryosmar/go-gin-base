package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harryosmar/go-gin-base/user"
	db "github.com/harryosmar/go-gin-base/user/DB"
)

func main() {

	r := gin.Default()

	database := db.ConnectDatabase()

	sqlDB, err := database.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDB.Close()

	userHandler := user.NewUserHandler(user.NewUserServiceMySQL(database))

	r.GET("/users", userHandler.List)
	r.GET("/user/:id", userHandler.Detail)
	r.PUT("/user/:id", userHandler.Update)
	r.POST("/user", userHandler.Create)
	r.DELETE("/user/:id", userHandler.Delete)
	r.Run(":8080")
}
