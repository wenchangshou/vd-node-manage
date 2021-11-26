package hashid

import (
	"errors"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/g"

	"github.com/speps/go-hashids/v2"
)

const (
	UserID = iota
	ProjectID
	FileID
	TaskID
	CategoryID
	DeviceID
)

var (
	ErrTypeNotMatch = errors.New("ID类型不匹配")
)

// HashEncode  对给定数据计算HashID
func HashEncode(v []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = g.Config().System.HashIDSalt
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	id, err := h.Encode(v)
	if err != nil {
		return "", err
	}
	return id, nil
}

func HashDecode(raw string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = g.Config().System.HashIDSalt
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return []int{}, err
	}
	return h.DecodeWithError(raw)
}

func HashID(id uint, t int) string {
	v, _ := HashEncode([]int{int(id), t})
	return v
}

// DecodeHashID 计算HashID对应的数据库ID
func DecodeHashID(id string, t int) (uint, error) {
	v, _ := HashDecode(id)
	if len(v) != 2 || v[1] != t {
		return 0, ErrTypeNotMatch
	}
	return uint(v[0]), nil
}
