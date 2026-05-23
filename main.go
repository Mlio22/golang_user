// @title Golang User Management API
// @version 1.0
// @description This is a simple user management API built with Go and Gin framework. It allows you to create, read, update, and delete users in memory. The API also includes Swagger documentation for easy testing and integration.
// @contact.name API Support
// @contact.url mailto:muhammatliopratama@gmail.com
// @contact.email muhammatliopratama@gmail.com
// @BasePath /
// @schemes http
package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "golang_user/docs"
	"golang_user/router"
)

func main() {
	r := gin.Default()
	router.UserRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
