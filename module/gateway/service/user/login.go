package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/common/util"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type LoginService struct {
	//TODO 细致调整验证规则
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}

func (service *LoginService) Login(c *gin.Context) serializer.Response {
	expectedUser, err := model.GetUserByUsername(service.UserName)
	if err != nil {
		return serializer.Err(serializer.CodeCredentialInvalid, "用户名和密码错误", err)
	}
	if authOk, _ := expectedUser.CheckPassword(service.Password); !authOk {
		return serializer.Err(serializer.CodeCredentialInvalid, "用户名和密码错误", nil)
	}
	if expectedUser.Status == model.Baned || expectedUser.Status == model.OveruseBaned {
		return serializer.Err(403, "该帐号未激活", err)
	}
	util.SetSession(c, map[string]interface{}{
		"user_id": expectedUser.ID,
	})
	return serializer.BuildUserResponse(expectedUser)
}
