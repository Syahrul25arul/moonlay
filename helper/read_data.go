package helper

import (
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ReadData(sheetname string) ([][]string, error) {
	f, err := excelize.OpenFile("dummies-migration-tests.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error open file")
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue(sheetname, "B2")
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error get cell value")
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error get rows")
	}
	return rows, nil
	// for _, row := range rows {
	// 	fmt.Print(row, "\t")
	// 	// for _, colCell := range row {
	// 	// 	fmt.Print(colCell, "\t")
	// 	// }
	// 	fmt.Println()
	// }
}
