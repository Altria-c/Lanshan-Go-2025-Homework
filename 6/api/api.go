package api

import (
	"net/http"

	"lansan-learning/6/dao"
	"lansan-learning/6/model"
	"lansan-learning/6/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	// 如果用户存在，这里这种是用户名可以一致的，即只要密码不一致就视为不同用户
	if dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}
	dao.AddUser(req.Username, req.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func Login(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 检查用户是否存在且密码是否正确
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	// 生成jwt token
	token, err := utils.MakeToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	//生成Refresh Token
	refreshToken, err := utils.MakeRefreshToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	// 返回token
	c.JSON(http.StatusOK, gin.H{
		"message":      "login success",
		"token":        token,
		"refreshtoken": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	type Refresh struct {
		RefreshToken string `json:"refreshtoken"`
	}
	var req Refresh
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	//
	username, err := utils.ParseToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid refresh token",
		})
		return
	}
	//
	NewToken, err := utils.MakeToken(username.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "generate token failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   NewToken,
	})
}

// 修改密码
func ChangePwd(c *gin.Context) {
	var req model.ChangeUser
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}
	// 验证用户和旧密码是否匹配
	if !dao.FindUser(req.Username, req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username or old password incorrect",
		})
		return
	}
	//更新密码
	if !dao.ChangePassword(req.Username, req.NewPassword) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "change failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "password changed",
	})
}
