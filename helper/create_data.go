package helper

import (
	"moonlay/model/domain"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func CreateData(f *excelize.File, sheetName string, data interface{}, domain string, activeSell int, iteration int) {

	// check sheeet
	indexSheet := f.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = f.NewSheet(sheetName)
	}

	f.SetActiveSheet(indexSheet)

	rv := reflect.ValueOf(data)
	datamart := ChooseDomain(domain)

	if iteration > 0 {
		activeSell++
	}
	for i := 0; i < rv.Len(); i++ {
		data := rv.Index(i)
		datamart = datamart.ChangeReflectValueToDataMart1(data)

		if activeSell == 1 {
			field := GetField(*datamart)
			f.SetSheetRow(sheetName, "A1", &field)
			activeSell++
		}

		dataRow := ConvertToString(*datamart)
		f.SetSheetRow(sheetName, "A"+strconv.Itoa(activeSell), &dataRow)

		activeSell++
	}

	if err := f.Save(); err != nil {
		PanicIFError(err)
	}
}

func CreateDataMart1(f *excelize.File, sheetName string, data []domain.Datamart1, activeSell int, iteration int) {
	// check sheeet
	indexSheet := f.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = f.NewSheet(sheetName)
	}

	f.SetActiveSheet(indexSheet)

	if iteration > 0 {
		activeSell++
	}

	for _, value := range data {
		datamart := &value

		if activeSell == 1 {
			field := GetField(*datamart)
			f.SetSheetRow(sheetName, "A1", &field)
			activeSell++
		}

		dataRow := ConvertToString(*datamart)
		f.SetSheetRow(sheetName, "A"+strconv.Itoa(activeSell), &dataRow)
		activeSell++
	}

	if err := f.Save(); err != nil {
		PanicIFError(err)
	}
}

func CreateDataMart2(f *excelize.File, sheetName string, data []domain.Datamart2, activeSell int, iteration int) {
	// check sheeet
	indexSheet := f.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = f.NewSheet(sheetName)
	}

	f.SetActiveSheet(indexSheet)

	if iteration > 0 {
		activeSell++
	}

	for _, value := range data {
		datamart := &value

		if activeSell == 1 {
			field := GetField(*datamart)
			f.SetSheetRow(sheetName, "A1", &field)
			activeSell++
		}

		dataRow := ConvertToString(*datamart)
		f.SetSheetRow(sheetName, "A"+strconv.Itoa(activeSell), &dataRow)
		activeSell++
	}

	if err := f.Save(); err != nil {
		PanicIFError(err)
	}
}

func CreateDataMart3(f *excelize.File, sheetName string, data []domain.Datamart3, activeSell int, iteration int) {
	// check sheeet
	indexSheet := f.GetSheetIndex(sheetName)
	if indexSheet == -1 {
		indexSheet = f.NewSheet(sheetName)
	}

	f.SetActiveSheet(indexSheet)

	if iteration > 0 {
		activeSell++
	}

	for _, value := range data {
		datamart := &value

		if activeSell == 1 {
			field := GetField(*datamart)
			f.SetSheetRow(sheetName, "A1", &field)
			activeSell++
		}

		dataRow := ConvertToString(*datamart)
		f.SetSheetRow(sheetName, "A"+strconv.Itoa(activeSell), &dataRow)
		activeSell++
	}

	if err := f.Save(); err != nil {
		PanicIFError(err)
	}
}
