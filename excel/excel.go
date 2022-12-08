package excel

import (
	"moonlay/model/domain"
)

type Excel interface {
	GetFile(name string) Excel
	CheckFile(nameFile string) bool
	CreateDataMart1(sheetName string, data []domain.Datamart1, activeSell int, iteration int)
	CreateDataMart2(sheetName string, data []domain.Datamart2, activeSell int, iteration int)
	CreateDataMart3(sheetName string, data []domain.Datamart3, activeSell int, iteration int)
	ReadData(sheetname, filename string) ([][]string, error)
}
