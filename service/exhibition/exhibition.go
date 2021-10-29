package exhibition

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
)

type GetComputerExhibitionService struct {
	ID string `uri:"id" form:"id"`
}

type WindowData struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	ID    string `json:"id"`
}
type ComputerExhibtionListData struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Level int          `json:"level"`
	Items []WindowData `json:"items"`
}
type GetComputerExhibitionFormData []ComputerExhibtionListData

func getResourceCategory(category string) string {
	resourceType := []string{
		"ppt", "video", "image", "pdf",
	}
	projectType := []string{
		"ue4", "app",
	}
	for _, t := range resourceType {
		if strings.EqualFold(t, category) {
			return "resource"
		}
	}
	for _, t := range projectType {
		if strings.EqualFold(category, t) {
			return "project"
		}
	}
	return "unknown"
}

func (service GetComputerExhibitionService) Get(c *gin.Context) serializer.Response {
	result := GetComputerExhibitionFormData{}
	categorys, err := model.GetComputerExhibtionCatetory(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取展项类别失败", err)
	}
	for _, category := range categorys {
		_data := &ComputerExhibtionListData{
			Name:  category.Name,
			ID:    category.ID,
			Level: category.Level,
		}
		exhibitions, _ := model.GetExhibitionByCategory(category.ID)
		for _, exhibition := range exhibitions {
			exhibitionData := &WindowData{
				ID:    exhibition.ID,
				Level: exhibition.Level,
				Name:  exhibition.Name,
			}
			_data.Items = append(_data.Items, *exhibitionData)
		}
		result = append(result, *_data)
	}

	return serializer.Response{
		Data: result,
	}
}

type ExhibitionService struct {
	ID         string `form:"id" json:"id" uri:"id"`
	CategoryID string `form:"category_id" json:"category_id" `
	Name       string `form:"name" json:"name" `
	Encryption bool   `form:"encryption" json:"encryption"`
	Password   string `form:"password" json:"password"`
	Control    string `form:"control" json:"control"`
	Windows    []struct {
		X         int      `json:"x"`
		Y         int      `json:"y"`
		Width     int      `json:"width"`
		Height    int      `json:"height"`
		Category  string   `json:"category"`
		Resources []string `json:"resources"`
		Active    string   `json:"active"`
	} `form:"windows" json:"windows"`
}

func (service *ExhibitionService) Create(c *gin.Context) serializer.Response {
	exhibition := model.Exhibition{
		CategoryID: service.CategoryID,
		Name:       service.Name,
		Encryption: service.Encryption,
		Password:   service.Password,
		Control:    service.Control,
	}
	tx := model.DB.Begin()
	result := tx.Model(&model.Exhibition{}).Create(&exhibition)
	if result.Error != nil {
		tx.Rollback()
		return serializer.Err(serializer.CodeDBError, "创建展项记录失败", result.Error)
	}
	for _, item := range service.Windows {
		w := &model.Window{
			Width:        item.Width,
			Height:       item.Height,
			X:            item.X,
			Y:            item.Y,
			Category:     item.Category,
			ExhibitionID: exhibition.ID,
		}
		if tx.Model(&model.Window{}).Create(&w).Error != nil {
			tx.Rollback()
			return serializer.Err(serializer.CodeDBError, "创建展项窗口失败，已经回滚", tx.Error)
		}
		for _, resource := range item.Resources {
			windowItem := &model.ExhibitionWindowItem{
				WindowID:     w.ID,
				ExhibitionID: exhibition.ID,
			}
			_category := getResourceCategory(item.Category)
			if _category == "project" {
				windowItem.ProjectID = resource
			} else if _category == "resource" {
				windowItem.ResourceID = resource
			}
			err := tx.Model(&model.ExhibitionWindowItem{}).Create(&windowItem).Error
			if err != nil {
				tx.Rollback()
				return serializer.Err(serializer.CodeDBError, "添加窗口资源项失败", err)
			}
			if item.Active == resource {
				tx.Model(&model.Window{}).Where("id=?", w.ID).Update("active_id", windowItem.ID)
			}

		}
	}
	tx.Commit()
	return serializer.Response{
		Data: exhibition.ID,
	}
}
func (service *ExhibitionService) Update() serializer.Response {
	exhibtion, err := model.GetExhibitionByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "没有找到对应的展项", err)
	}
	err = model.DeleteExhibtionWindowByExhibitionID(exhibtion.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除子窗口失败", err)
	}
	maps := map[string]interface{}{
		"name":       service.Name,
		"encryption": service.Encryption,
		"password":   service.Password,
		"control":    service.Control,
	}
	err = exhibtion.Update(maps)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "更新展项失败", err)
	}
	// for _, item := range service.Items {
	// 	_item := &model.Window{
	// 		Width:        item.Width,
	// 		Height:       item.Height,
	// 		X:            item.X,
	// 		Y:            item.Y,
	// 		Category:     item.Category,
	// 		ExhibitionID: exhibtion.ID,
	// 	}
	// 	_item.Create()
	// }
	query := GetExhibitionDetailsService{
		ID: exhibtion.ID,
	}
	return query.Get()
}
func (service *ExhibitionService) Delete() serializer.Response {
	err := model.DeleteExhibitionWindowItemByExhibitionID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除展项窗口资源失败", err)
	}
	err = model.DeleteExhibtionWindowByExhibitionID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除展项窗口失败", err)
	}
	err = model.DeleteExhibtionByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "删除展项失败", err)
	}
	return serializer.Response{}
}

type GetExhibitionDetailsService struct {
	ID string `json:"id" uri:"id" form:"id"`
}
type ResourceInfo struct {
	ID         string `json:"id"`
	ResourceID string `json:"resource_id"`
	Name       string `json:"name"`
}
type WindowInfo struct {
	ID       string `json:"id"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Category string `json:"category"`
	Active   string `json:"active"`
	Resource []*ResourceInfo
}
type GetExhibtionDetailsResultData struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Level      int           `json:"level"`
	Encryption bool          `json:"encryption"`
	Windows    []*WindowInfo `json:"windows"`
}

func (service *GetExhibtionDetailsResultData) GetWindow(id string) *WindowInfo {
	for _, window := range service.Windows {
		if window.ID == id {
			return window
		}
	}
	return nil
}

func (service *GetExhibitionDetailsService) Get() serializer.Response {
	result := &GetExhibtionDetailsResultData{}
	exhibition, err := model.GetExhibitionByID(service.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取展项失败", err)
	}
	items, err := model.GetExhibtionDetailsByExhibtionID(exhibition.ID)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取展项详细信息失败", err)
	}
	result.ID = exhibition.ID
	result.Name = exhibition.Name
	result.Level = exhibition.Level
	result.Encryption = exhibition.Encryption

	for _, item := range items {
		var window *WindowInfo
		resource := &ResourceInfo{}
		window = result.GetWindow(item.WindowID)
		if window == nil {
			window = &WindowInfo{
				ID:       item.WindowID,
				X:        item.Window.X,
				Y:        item.Window.Y,
				Width:    item.Window.Width,
				Height:   item.Window.Height,
				Category: item.Window.Category,
				Active:   item.Window.ActiveID,
			}
			result.Windows = append(result.Windows, window)
		}
		category := getResourceCategory(window.Category)
		if category == "project" {
			resource.Name = item.Project.Name
			resource.ID = item.ID
			resource.ResourceID = item.ProjectID
		} else if category == "resource" {
			resource.Name = item.Resource.Name
			resource.ID = item.ID
			resource.ResourceID = item.ResourceID

		}
		window.Resource = append(window.Resource, resource)

	}
	return serializer.Response{
		Data: result,
	}
	// items, err := model.GetExhibitionItemsByExhibitionID(exhibition.ID)
	// if err != nil {
	// 	return serializer.Err(serializer.CodeDBError, "获取窗口失败", err)
	// }
	// result.Name = exhibition.Name
	// result.Level = exhibition.Level
	// result.Encryption = exhibition.Encryption
	// result.Items = make([]WindowInfo, 0)
	// for _, item := range items {
	// 	window := &WindowInfo{
	// 		X:      item.X,
	// 		Y:      item.Y,
	// 		Width:  item.Width,
	// 		Height: item.Height,
	// 	}
	// 	if item.Category == "project" {
	// 		ids := make([]string, 0)
	// 		ids = append(ids, item.ProjectID)
	// 		projects, err := model.GetProjectByIds(ids)
	// 		if len(projects) == 0 || err != nil {
	// 			return serializer.Err(serializer.CodeDBError, "获取对应的项目失败", err)
	// 		}
	// 		window.Category = projects[0].Category
	// 		window.Name = projects[0].Name
	// 	} else if item.Category == "resource" {
	// 		resource, err := model.GetResourceById(item.ResourceID)
	// 		if resource.ID == "" || err != nil {
	// 			return serializer.Err(serializer.CodeDBError, "获取项目资源项目", err)
	// 		}
	// 		window.Category = resource.Category
	// 		window.Name = resource.Name
	// 	}
	// 	result.Items = append(result.Items, *window)
	// }

}

// type UpdateExhibtionService struct {
// 	ID         string `json:"id" form:"id"`
// 	Name       string `json:"name" form:"name"`
// 	Encryption bool   `json:"encryption" form:"encrytpion"`
// 	Password   string `json:"password" form:"password"`
// 	Items      []struct {
// 		X          int
// 		Y          int
// 		Width      int
// 		Height     int
// 		Category   string
// 		ResourceID string `json:"resource_id" form:"resource_id"`
// 		ProjectID  string `json:"project_id" form:"project_id"`
// 	} `form:"items" json:"items"`
// }
