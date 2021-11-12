package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetActiveUserByID(uid.(string))
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
