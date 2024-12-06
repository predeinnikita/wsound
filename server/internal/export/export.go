package export

import (
	"animal-sound-recognizer/internal/audio"
	"animal-sound-recognizer/internal/projects"
	"fmt"
	"github.com/xuri/excelize/v2"
	"unicode/utf8"
)

func createExcelDataByProjectId(projectId uint64) ([]byte, string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	project, _ := projects.GetProject(projectId)
	addProjectInfo(f, project)
	f.DeleteSheet("Sheet1")

	fileName := fmt.Sprintf("%s.xlsx", project.Name)
	buffer, _ := f.WriteToBuffer()

	return buffer.Bytes(), fileName
}

func createExcelAllData() ([]byte, string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	projects, _ := projects.GetAllProjects()
	for _, project := range projects {
		addProjectInfo(f, project)
	}
	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(0)

	fileName := "data.xlsx"
	buffer, _ := f.WriteToBuffer()

	return buffer.Bytes(), fileName
}

func addProjectInfo(f *excelize.File, project projects.Project) {
	audios, _ := audio.GetAllAudios(project.ID)
	offset := 0
	index, _ := f.NewSheet(project.Name)
	f.SetActiveSheet(index)

	style, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"BDD7EE"},
			Pattern: 1,
		},
	})

	f.SetCellValue(project.Name, "A1", "Название аудио")
	f.SetCellValue(project.Name, "B1", "Статус")
	f.SetCellValue(project.Name, "C1", "Интервалы с воем")
	f.SetCellStyle(project.Name, "A1", "C1", style)

	for i, audio := range audios {
		if len(audio.Intervals) > 1 {
			f.MergeCell(
				project.Name,
				fmt.Sprintf("A%d", i+offset+2),
				fmt.Sprintf("A%d", i+len(audio.Intervals)+offset+1),
			)

			f.MergeCell(
				project.Name,
				fmt.Sprintf("B%d", i+offset+2),
				fmt.Sprintf("B%d", i+len(audio.Intervals)+offset+1),
			)
		}

		f.SetCellValue(
			project.Name,
			fmt.Sprintf("A%d", i+offset+2),
			audio.Name,
		)

		f.SetCellValue(
			project.Name,
			fmt.Sprintf("B%d", i+offset+2),
			audio.Status,
		)

		for j, interval := range audio.Intervals {
			f.SetCellValue(
				project.Name,
				fmt.Sprintf("C%d", i+j+2+offset),
				fmt.Sprintf("%s - %s", interval.Start, interval.End),
			)
		}

		offset += len(audio.Intervals) - 1
	}

	fitColumns(f, project.Name)
}

func fitColumns(f *excelize.File, sheetName string) {
	cols, _ := f.GetCols(sheetName)

	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, _ := excelize.ColumnNumberToName(idx + 1)
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}
