package excel

import (
	"errors"
	"fmt"
	"moonlay/helper"
	"moonlay/model/domain"
	"os"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type excelImpl struct {
	file *excelize.File
}

func NewExcel() Excel {
	return &excelImpl{
		file: &excelize.File{},
	}
}

// * Get file excel, if file not exists create new file
func (e *excelImpl) GetFile(name string) Excel {
	file, err := excelize.OpenFile(name)
	if err != nil {
		file := excelize.NewFile()

		file.SetSheetName("Sheet1", "datamart")
		file.NewSheet("datamart1")
		file.NewSheet("datamart2")
		err := file.SaveAs(name)
		helper.PanicIFError(err)
	}

	errs := e.CheckFile(name)
	if errs {
		helper.PanicIFError(errors.New("file not exists"))
	}

	return &excelImpl{
		file: file,
	}
}

func (e *excelImpl) CheckFile(nameFile string) bool {
	// define an interval and the ticker for this interval
	interval := 2 * time.Second
	// create a new Ticker
	tk := time.NewTicker(interval)

	var fileExist bool

	// start the ticker by constructing a loop
	i := 0
	for range tk.C {
		i++

		if i == 5 {
			fileExist = true
			break
		}
		_, err := os.Stat(nameFile)
		if err == nil {
			fileExist = false
			break
		}
	}
	tk.Stop()
	return fileExist
}
func (e *excelImpl) CreateDataMart1(sheetName string, data []domain.Datamart1, activeSell int, iteration int) {
	// check sheeet
	indexSheet := e.file.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = e.file.NewSheet(sheetName)
	}

	e.file.SetActiveSheet(indexSheet)

	// * set iteration to 1 for cell in excel
	if iteration > 0 {
		activeSell++
	}

	// TODO: set struct and set row for field in excel
	datamart := &domain.Datamart1{}
	setFieldRow(e.file, sheetName, activeSell, func() []string {
		return helper.GetField(*datamart)
	})

	// TODO: set data to excel
	setDataToExcel(e.file, sheetName, len(data), activeSell, func(i int) []string {
		datamart = &data[i]
		return helper.ConvertToString(*datamart)
	})

	// TODO: save file if done
	if err := e.file.Save(); err != nil {
		helper.PanicIFError(err)
	}
}

func (e *excelImpl) CreateDataMart2(sheetName string, data []domain.Datamart2, activeSell int, iteration int) {
	// check sheeet
	indexSheet := e.file.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = e.file.NewSheet(sheetName)
	}

	e.file.SetActiveSheet(indexSheet)

	// * set iteration to 1 for cell in excel
	if iteration > 0 {
		activeSell++
	}

	// TODO: set struct and set row for field in excel
	datamart := &domain.Datamart2{}
	setFieldRow(e.file, sheetName, activeSell, func() []string {
		return helper.GetField(*datamart)
	})

	// TODO: set data to excel
	setDataToExcel(e.file, sheetName, len(data), activeSell, func(i int) []string {
		datamart = &data[i]
		return helper.ConvertToString(*datamart)
	})

	// TODO: save file if done
	if err := e.file.Save(); err != nil {
		helper.PanicIFError(err)
	}
}

func (e *excelImpl) CreateDataMart3(sheetName string, data []domain.Datamart3, activeSell int, iteration int) {
	// check sheeet
	indexSheet := e.file.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = e.file.NewSheet(sheetName)
	}

	e.file.SetActiveSheet(indexSheet)

	// * set iteration to 1 for cell in excel
	if iteration > 0 {
		activeSell++
	}

	// TODO: set struct and set row for field in excel
	datamart := &domain.Datamart3{}
	setFieldRow(e.file, sheetName, activeSell, func() []string {
		return helper.GetField(*datamart)
	})

	// TODO: set data to excel
	setDataToExcel(e.file, sheetName, len(data), activeSell, func(i int) []string {
		datamart = &data[i]
		return helper.ConvertToString(*datamart)
	})

	// TODO: save file if done
	if err := e.file.Save(); err != nil {
		helper.PanicIFError(err)
	}
}

func (e *excelImpl) ReadData(sheetname, filename string) ([][]string, error) {
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

func setDataToExcel(file *excelize.File, sheetName string, len int, activeSell int, callback func(i int) []string) {
	for i := 0; i < len; i++ {
		dataRow := callback(i)
		file.SetSheetRow(sheetName, "A"+strconv.Itoa(activeSell), &dataRow)
		activeSell++
	}
}

func setFieldRow(file *excelize.File, sheetName string, activeSell int, callback func() []string) {
	if activeSell == 1 {
		field := callback()
		file.SetSheetRow(sheetName, "A1", &field)
		activeSell++
	}
}
