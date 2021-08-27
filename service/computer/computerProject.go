package computer

import (
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/serializer"
	"github.com/wenchangshou2/vd-node-manage/pkg/util"
)

type ApplicationParams struct {
	StartPath     string   `json:"start_path"`
	DelayStart    int      `json:"delay_start"`
	ParameterMode string   `json:"parameter_mode"`
	Parameters    string   `json:"parameters"`
	Type          string   `json:"type"`
	Key           []string `json:"key"`
	Message       []string `json:"message"`
}
type ApplicationItem struct {
	ID        int               `json:"id"`
	Name      string            `json:"name"`
	Path      string            `json:"path"`
	Arguments ApplicationParams `json:"arguments"`
}
type ApplicationList []ApplicationItem

func (service *ApplicationList) Add(projectRelease model.ProjectRelease) {
	item := ApplicationItem{
		ID:   int(projectRelease.ID),
		Name: projectRelease.Project.Name,
		Path: projectRelease.File.Uuid,
		Arguments: ApplicationParams{
			StartPath:     projectRelease.Project.Start,
			DelayStart:    0,
			ParameterMode: "append",
			Parameters:    "",
			Type:          projectRelease.Project.Category,
			Key:           make([]string, 0),
			Message:       make([]string, 0),
		},
	}
	*service = append((*service), item)
}

type ResourceList struct {
	Image []ResourceItem `json:"image"`
	Pdf   []ResourceItem `json:"pdf"`
	Ppt   []ResourceItem `json:"ppt"`
	Video []ResourceItem `json:"video"`
}

func (service *ResourceList) Add(resource *model.Resource) {

	item := ResourceItem{
		Name: resource.Name,
		Path: resource.File.SourceName,
		ID:   int(resource.ID),
	}
	switch resource.Category {
	case "image":
		service.Image = append(service.Image, item)
	case "video":
		service.Video = append(service.Video, item)
	case "ppt":
		service.Ppt = append(service.Ppt, item)
	case "pdf":
		service.Pdf = append(service.Pdf, item)
	}
}

func NewResourceList() ResourceList {
	return ResourceList{
		Image: make([]ResourceItem, 0),
		Pdf:   make([]ResourceItem, 0),
		Ppt:   make([]ResourceItem, 0),
		Video: make([]ResourceItem, 0),
	}
}

type ResourceItem struct {
	Path string `json:"path"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type ComputerProjectGetCrossResource struct {
}

// CrossResourceForm 交叉资源表单
type CrossResourceForm struct {
	Application ApplicationList `json:"application"`
	Resource    ResourceList    `json:"resource"`
}

func (service *ComputerProjectGetCrossResource) getProject(ids []int) (ApplicationList, error) {
	arrayList := ApplicationList{}
	projectReleaseMap := make(map[uint][]int)
	computerProjectList, err := model.GetComputerCrossProject(ids)
	if err != nil {
		return nil, err
	}
	for _, cp := range computerProjectList {
		if projectReleaseItem, ok := projectReleaseMap[cp.ProjectReleaseID]; ok {
			projectReleaseMap[cp.ProjectReleaseID] = append(projectReleaseItem, int(cp.ComputerId))
			continue
		}
		projectReleaseMap[cp.ProjectReleaseID] = make([]int, 0)
		projectReleaseMap[cp.ProjectReleaseID] = append(projectReleaseMap[cp.ProjectReleaseID], int(cp.ComputerId))
	}
	for projectReleaseID, computerIds := range projectReleaseMap {
		if util.IntArrayEquals(ids, computerIds) {
			pr, err := model.GetProjectReleaseByID(projectReleaseID)
			if err != nil {
				continue
			}
			arrayList.Add(pr)
		}
	}
	return arrayList, nil
}

func (service *ComputerProjectGetCrossResource) getResource(ids []int) (ResourceList, error) {
	resourceMap := make(map[uint][]int)
	resourceList := NewResourceList()
	// 获取资源
	computerResources, err := model.GetComputerResourcesByComputerIds(ids)
	if err != nil {
		return resourceList, err
	}
	for _, cr := range computerResources {
		if _, ok := resourceMap[cr.ResourceID]; ok {
			resourceMap[cr.ResourceID] = append(resourceMap[cr.ResourceID], int(cr.ComputerID))
			continue
		}
		resourceMap[cr.ResourceID] = make([]int, 0)
		resourceMap[cr.ResourceID] = append(resourceMap[cr.ResourceID], int(cr.ComputerID))
	}
	for resourceID, computerIds := range resourceMap {
		if util.IntArrayEquals(computerIds, ids) {
			resource, err := model.GetResourceById(resourceID)
			if err != nil {
				continue
			}
			resourceList.Add(resource)
		}
	}
	return resourceList, nil
}

// Get 获取交叉资源
func (service *ComputerProjectGetCrossResource) Get() serializer.Response {
	var (
		err error
	)
	computers, _ := model.ListComputer()
	form := CrossResourceForm{}

	ids := make([]int, 0)
	for _, computer := range computers {
		ids = append(ids, int(computer.ID))
	}
	form.Application, err = service.getProject(ids)
	if err != nil {
		serializer.Err(serializer.CodeDBError, "获取计算机项目失败", nil)
	}
	form.Resource, err = service.getResource(ids)
	if err != nil {
		serializer.Err(serializer.CodeDBError, "获取计算机资源列表失败", err)
	}
	return serializer.Response{
		Data: form,
	}
}

type ComputerProjectListService struct {
	ID int `json:"id" uri:"id"`
}
type ComputerProjectListForm struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Arguments   string `json:"arguments"`
	Control     string `json:"control"`
	CoverUri    string `json:"cover_uri"`
}

func (service *ComputerProjectListService) List() serializer.Response {
	computerProject, err := model.GetComputerProjectByComputerId(service.ID)
	ids := make([]int, 0)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取计算机项目失败", err)
	}
	for _, cp := range computerProject {
		ids = append(ids, int(cp.ProjectID))
	}
	projects, err := model.GetProjectByIds(ids)
	if err != nil {
		return serializer.Err(serializer.CodeDBError, "获取项目列表失败", err)
	}
	return serializer.Response{
		Data: projects,
	}
}