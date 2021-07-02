package model

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/wenchangshou2/vd-node-manage/pkg/util"
	"gorm.io/gorm"
)

const (
	// Active 账户正常状态
	Active = iota
	// NotActivicated 未激活
	NotActivicated
	// Baned 被封禁
	Baned
	// OveruseBaned 超额使用被封禁
	OveruseBaned
)

type User struct {
	// 表字段
	gorm.Model
	Username string `gorm:"size:50"`
	Password string `json:"-"`
	Status   int
	Type     int
}

func GetUserByUsername(username string) (User, error) {
	var user User
	result := DB.Debug().Where("username=?", username).First(&user)
	return user, result.Error
}

//GetActiveUserByID 通过id获取用户
func GetActiveUserByID(ID interface{}) (User, error) {
	var user User
	result := DB.Where("status=?", Active).First(&user, ID)
	return user, result.Error
}
func (user *User) CheckPassword(password string) (bool, error) {
	passwordStore := strings.Split(user.Password, ":")
	if len(passwordStore) != 2 && len(passwordStore) != 3 {
		return false, errors.New("unknown password type")
	}
	if len(passwordStore) == 3 {
		if passwordStore[0] != "md5" {
			return false, errors.New("unknown password type")
		}
		hash := md5.New()
		_, err := hash.Write([]byte(passwordStore[2] + password))
		bs := hex.EncodeToString(hash.Sum(nil))
		if err != nil {
			return false, err
		}
		return bs == passwordStore[1], nil
	}
	hash := sha1.New()
	_, err := hash.Write([]byte(password + passwordStore[0]))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return false, err
	}
	return bs == passwordStore[1], nil
}
func (user *User) SetPassword(password string) error {
	salt := util.RandStringRunes(16)

	hash := sha1.New()
	_, err := hash.Write([]byte(password + salt))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return err
	}
	user.Password = salt + ":" + string(bs)
	return nil
}

func NewUser() User {
	return User{}
}
