package main

import (
	"auth-api/m/controllers"
	"auth-api/m/initializers"
	"auth-api/m/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middleware.CheckAuth, controllers.GetUserProfile)
	router.Run()
}
