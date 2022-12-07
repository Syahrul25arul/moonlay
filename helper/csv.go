package helper

import (
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

func GetFileCsv(name string) *excelize.File {
	var file *excelize.File
	file, err := excelize.OpenFile(name)
	if err != nil {
		file := excelize.NewFile()

		file.SetSheetName("Sheet1", "datamart")
		file.NewSheet("datamart1")
		file.NewSheet("datamart2")
		err := file.SaveAs(name)
		PanicIFError(err)
	}

	return file
}

func CheckFile(nameFile string) (*excelize.File, bool) {

	// define an interval and the ticker for this interval
	interval := 2 * time.Second
	// create a new Ticker
	tk := time.NewTicker(interval)

	var fileExist bool
	var file *excelize.File

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
			file = GetFileCsv(nameFile)
			break
		}
	}
	tk.Stop()
	return file, fileExist
}
