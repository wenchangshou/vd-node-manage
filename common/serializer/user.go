package serializer

import (
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"time"
)

func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// User 用户序列化器
type User struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UserName:  user.Username,
	}

}
