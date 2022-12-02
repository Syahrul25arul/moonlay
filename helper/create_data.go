package helper

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func CreateData(data [][]string) {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("datamart_1")
	// Set value of a cell.
	iLoopRow := 1
	for _, value := range data {
		var cell string
		iLoopCell := 1

		for _, colCell := range value {

			switch {
			case iLoopCell == 1:
				cell = "A"
			case iLoopCell == 2:
				cell = "B"
			case iLoopCell == 3:
				cell = "C"
			case iLoopCell == 4:
				cell = "D"
			}

			// Set value of a cell.
			f.SetCellValue("datamart_1", cell+strconv.Itoa(iLoopRow), colCell)
			// Set active sheet of the workbook.
			f.SetActiveSheet(index)

			if iLoopCell == 4 {
				iLoopCell = 1
				iLoopRow++
			}

			iLoopCell++
		}
	}

	fmt.Println(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("datamart1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
