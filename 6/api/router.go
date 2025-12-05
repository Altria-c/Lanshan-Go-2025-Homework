package api

import (
	"lansan-learning/6/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter_gin() {
	r := gin.Default()
	r.GET("/login")
	r.POST("login", Login)

	r.GET("/register")
	r.POST("register", Register)

	r.GET("/refresh")
	r.POST("refresh", RefreshToken)

	r.GET("/changepwd")
	r.POST("changepwd", middleware.JWTAuth(), ChangePwd)

	r.Run(":9090")
}
