package file

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type FileCreateService struct {
	Name string `json:"name" form:"name"`
	Mode string `json:"mode" form:"mode" binding:"required"`
	// SourceName string `json:"source_name" `
	// UserId     uint   `json:"user_id"`
	// Size       uint   `json:"size"`
}

func (service *FileCreateService) Create(c *gin.Context, user *model.User) serializer.Response {
	var (
		id uint
	)
	// 如果当前采用本地上传
	if service.Mode == "upload" {
		_file, err := c.FormFile("file")
		if err != nil {
			return serializer.Err(serializer.CodeUploadFailed, "上错文件失败", err)
		}
		ext := filepath.Ext(_file.Filename)
		newFileName := uuid.New().String() + ext
		sourceName := "upload/" + newFileName
		if err := c.SaveUploadedFile(_file, sourceName); err != nil {
			return serializer.Err(serializer.CodeUploadFailed, "保存上传文件失败", err)
		}
		fileModel := model.File{
			Name:       _file.Filename,
			SourceName: sourceName,
			Mode:       service.Mode,
			Size:       uint(_file.Size),
			UserId:     user.ID,
		}
		id, err = fileModel.Create()
		if err != nil {
			return serializer.Err(serializer.CodeDBError, "创建文件记录失败", err)
		}
		return serializer.Response{
			Data: id,
		}
	}
	return serializer.Response{}
}
