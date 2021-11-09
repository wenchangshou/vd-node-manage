package controllers

import (
	"errors"
	user2 "github.com/wenchangshou2/vd-node-manage/module/gateway/service/user"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var service user2.LoginService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func GetCurrentUser(c *gin.Context) {
	user := CurrentUser(c)
	if user == nil {
		c.JSON(200, ErrorResponse(errors.New("未登录")))
		return
	}
	c.JSON(200, user)
}
func UserRegister(c *gin.Context) {
	var service user2.RegisterService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
