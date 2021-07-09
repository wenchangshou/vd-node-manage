package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/service/user"
)

func UserLogin(c *gin.Context) {
	var service user.UserLoginService
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
	var service user.UserRegisterService
	if err := c.ShouldBindJSON(&service); err == nil {
		res := service.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
