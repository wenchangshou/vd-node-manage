package model

// type ComputerProject struct {
// 	Base
// 	ComputerId       string `gorm:"computer_id"`
// 	ProjectID        string
// 	ProjectReleaseID string
// 	Status           uint `gorm:"status"`
// 	ProjectRelease   ProjectRelease
// }

// func (cp *ComputerProject) Create() (string, error) {
// 	if err := DB.Create(cp).Error; err != nil {
// 		return "", err
// 	}
// 	if cp.ID == "" {
// 		return "", errors.New("未找到相应的记录")
// 	}
// 	return cp.ID, nil
// }
// func GetComputerProjectByID(cid string, pid string) (ComputerProject, error) {
// 	var computerProject ComputerProject
// 	result := DB.Debug().Model(&ComputerProject{}).Where("computer_id=? AND project_id  = ?", cid, pid).First(&computerProject)
// 	return computerProject, result.Error
// }

// func GetComputerProjectByProjectIDAndProjectReleaseID(projectID, projectReleaseID string) ([]ComputerProject, error) {
// 	var computerProject []ComputerProject
// 	result := DB.Debug().Where("project_id = ? and project_release_id = ?", projectID, projectReleaseID).Find(&computerProject)
// 	return computerProject, result.Error
// }
// func GetComputerCrossProject(ids []string) ([]ComputerProject, error) {
// 	var computerProject []ComputerProject
// 	result := DB.Debug().Model(&ComputerProject{}).Where("computer_id IN ?", ids).Find(&computerProject)
// 	return computerProject, result.Error
// }

// // GetComputerProjectReleaseByComputerID 获取指定计算机的资源
// func GetComputerProjectReleaseByComputerID(id string) ([]ProjectRelease, error) {
// 	var computerProjectList []ComputerProject
// 	projectReleaseList := make([]ProjectRelease, 0)
// 	result := DB.Debug().Model(&ComputerProject{}).Find(&computerProjectList)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	for _, computerProject := range computerProjectList {
// 		projeectRelease, err := GetProjectReleaseByIdAndProjectId(computerProject.ProjectID, computerProject.ProjectReleaseID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		projectReleaseList = append(projectReleaseList, projeectRelease)
// 	}
// 	return projectReleaseList, nil
// }
// func DeleteComputerProjectByID(id int) error {
// 	result := DB.Debug().Delete(&ComputerProject{}, id)
// 	return result.Error
// }

// // 通过项目id列表来批量获取
// func GetComputerProjectByProjectIds(ids []string) ([]ComputerProject, error) {
// 	var computerProjectList []ComputerProject
// 	result := DB.Debug().Where("project_id IN ?", ids).Find(&computerProjectList)
// 	return computerProjectList, result.Error
// }

// // GetComputerProjectByProjectID 获取计算机项目通过项目id
// func GetComputerProjectByProjectID(id int) ([]ComputerProject, error) {
// 	var computerProjectList []ComputerProject
// 	result := DB.Where("project_id = ? ", id).Find(&computerProjectList)
// 	return computerProjectList, result.Error
// }
// func ListComputerProject() ([]ComputerProject, error) {
// 	var computerProjectList []ComputerProject
// 	result := DB.Model(&ComputerProject{}).Find(&computerProjectList)
// 	return computerProjectList, result.Error
// }
// func GetComputerProjectByComputerId(id string) ([]ComputerProject, error) {
// 	var computerProjectList []ComputerProject
// 	result := DB.Model(&ComputerProject{}).Where("computer_id=?", id).Find(&computerProjectList)
// 	return computerProjectList, result.Error
// }
