package system

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx/v3"
	"github.com/wenchangshou2/vd-node-manage/model"
	"github.com/wenchangshou2/vd-node-manage/pkg/util"
)

type ProjectExportService struct {
}
type ProjectInformation struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	PublishDate time.Time `json:"publish_time`
}

// Export 导出项目成excel文件
func (projectExportService ProjectExportService) Export() (*xlsx.File, error) {
	computerNameMap := make(map[uint]string)
	computerProjectMap := make(map[uint][]ProjectInformation)
	computerProjectList, err := model.ListComputerProject()
	file := xlsx.NewFile()
	if err != nil {
		return nil, err
	}

	computers, _ := model.ListComputer()
	for _, computer := range computers {
		computerNameMap[computer.ID] = computer.Name
		computerProjectMap[computer.ID] = make([]ProjectInformation, 0)
	}
	for _, cp := range computerProjectList {
		projectRelease, err := model.GetProjectReleaseByID(cp.ProjectReleaseID)
		if err != nil {
			return nil, err
		}
		projectInformation := ProjectInformation{
			Name:        projectRelease.Project.Name,
			PublishDate: cp.CreatedAt,
			Version:     projectRelease.Tag,
		}

		computerProjectMap[cp.ComputerId] = append(computerProjectMap[cp.ComputerId], projectInformation)
	}
	t := make([]string, 0)
	t = append(t, "项目名称")
	t = append(t, "版本号")
	t = append(t, "添加时间")
	for computerID, projects := range computerProjectMap {
		computerName := computerNameMap[computerID]
		sheet, err := file.AddSheet(computerName)
		if err != nil {
			return nil, err
		}
		titleRow := sheet.AddRow()
		titleRow.SetHeight(100)
		xlsRow := util.NewRow(titleRow, t)
		err = xlsRow.SetRowTitle()
		if err != nil {
			return nil, err
		}
		for _, project := range projects {
			currentRow := sheet.AddRow()
			currentRow.SetHeight(100)
			tmp := make([]string, 0)
			tmp = append(tmp, string(project.Name))
			tmp = append(tmp, project.Version)
			_date := fmt.Sprintf("%d年%d月%d日 %d时%d分", project.PublishDate.Year(), project.PublishDate.Month(), project.PublishDate.Day(), project.PublishDate.Hour(), project.PublishDate.Minute())
			tmp = append(tmp, _date)
			xlsRow := util.NewRow(currentRow, tmp)
			err := xlsRow.GenerateRow()
			if err != nil {
				return nil, err
			}
		}
	}
	return file, nil
}
