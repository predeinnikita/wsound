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

	audios, _ := audio.GetAllAudios(projectId)
	project, _ := projects.GetProject(projectId)
	sheet := "Sheet1"

	style, _ := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"BDD7EE"},
			Pattern: 1,
		},
	})

	f.SetCellValue(sheet, "A1", "Название аудио")
	f.SetCellValue(sheet, "B1", "Статус")
	f.SetCellValue(sheet, "C1", "Интервалы с воем")
	f.SetCellStyle(sheet, "A1", "C1", style)

	offset := 0

	for i, audio := range audios {
		if len(audio.Intervals) > 1 {
			f.MergeCell(
				sheet,
				fmt.Sprintf("A%d", i+offset+2),
				fmt.Sprintf("A%d", i+len(audio.Intervals)+offset+1),
			)

			f.MergeCell(
				sheet,
				fmt.Sprintf("B%d", i+offset+2),
				fmt.Sprintf("B%d", i+len(audio.Intervals)+offset+1),
			)
		}

		f.SetCellValue(
			sheet,
			fmt.Sprintf("A%d", i+offset+2),
			audio.Name,
		)

		f.SetCellValue(
			sheet,
			fmt.Sprintf("B%d", i+offset+2),
			audio.Status,
		)

		for j, interval := range audio.Intervals {
			f.SetCellValue(
				sheet,
				fmt.Sprintf("C%d", i+j+2+offset),
				fmt.Sprintf("%s - %s", interval.Start, interval.End),
			)
		}

		offset += len(audio.Intervals) - 1
	}

	fitColumns(f, sheet)

	fileName := fmt.Sprintf("%s.xlsx", project.Name)
	buffer, _ := f.WriteToBuffer()

	return buffer.Bytes(), fileName
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
