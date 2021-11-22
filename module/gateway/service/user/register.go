package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/common/serializer"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
)

type RegisterService struct {
	//TODO 细致调整验证规则
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}

func (service *RegisterService) Register(c *gin.Context) serializer.Response {

	//创建新的用户对象
	user := model.NewUser()
	user.Username = service.UserName
	user.SetPassword(service.Password)
	user.Status = model.Active
	if err := model.DB.Create(&user).Error; err != nil {
		expectedUser, err := model.GetUserByUsername(service.UserName)
		if expectedUser.Status == model.NotActivated {
			user = expectedUser
		} else {
			return serializer.DBErr("此邮箱已存在", err)
		}
	}
	return serializer.Response{}
}