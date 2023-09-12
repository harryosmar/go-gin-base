package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harryosmar/go-gin-base/user"
)

func main() {
	r := gin.Default()

	userHandler := user.NewUserHandler(user.NewUserServiceMySQL())

	r.GET("/users", userHandler.List)
	r.GET("/user/:id", userHandler.Detail)
	r.PUT("/user/:id", userHandler.Update)
	r.POST("/user", userHandler.Create)
	r.DELETE("/user/:id", userHandler.Delete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
