package middleware

import (
	"lansan-learning/6/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// jwt验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求头为空",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求头格式有误",
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "无效的Token",
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()

	}

}
