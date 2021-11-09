package file

import (
	"github.com/wenchangshou2/vd-node-manage/module/gateway/model"
	"github.com/wenchangshou2/vd-node-manage/module/gateway/pkg/serializer"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wenchangshou2/zutil"
)

type CreateService struct {
	Name string `json:"name" form:"name"`
	Mode string `json:"mode" form:"mode" binding:"required"`
	// SourceName string `json:"source_name" `
	// UserId     uint   `json:"user_id"`
	// Size       uint   `json:"size"`
}

func (service *CreateService) Create(c *gin.Context, user *model.User) serializer.Response {
	var (
		id string
	)
	// 如果当前采用本地上传
	if service.Mode == "upload" {
		_file, err := c.FormFile("file")
		if err != nil {
			return serializer.Err(serializer.CodeUploadFailed, "上错文件失败", err)
		}
		ext := filepath.Ext(_file.Filename)
		uid := uuid.New().String()
		newFileName := uid + ext
		sourceName := newFileName
		if err := c.SaveUploadedFile(_file, "upload/"+sourceName); err != nil {
			return serializer.Err(serializer.CodeUploadFailed, "保存上传文件失败", err)
		}
		fileModel := model.File{
			Name:       _file.Filename,
			SourceName: sourceName,
			Mode:       service.Mode,
			Size:       uint(_file.Size),
			UserId:     user.ID,
			Uuid:       uid,
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

type FileDownloadService struct {
	ID string `uri:"id" form:"id"`
}

func (service FileDownloadService) Download(c *gin.Context) serializer.Response {
	file, err := model.GetFileById(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeNotFindFile, "没有找到文件", err)
	}
	filePath := path.Join("upload/", file.SourceName)
	if !zutil.IsExist(filePath) {
		return serializer.Err(serializer.CodeNotFindFile, "没有找到文件", nil)
	}
	fi, _ := os.Stat(filePath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+file.Name)
	c.Header("Content-Length", strconv.Itoa(int(fi.Size())))

	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
	return serializer.Response{}
}
