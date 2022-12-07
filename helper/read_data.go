package helper

import (
	"errors"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ReadData(sheetname, filename string) ([][]string, error) {
	f, err := excelize.OpenFile(filename)
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

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error get rows")
	}
	return rows, nil
}
