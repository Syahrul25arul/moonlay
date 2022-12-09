package helper

import (
	"reflect"
)

// ? get field from struct
func GetField(data interface{}) []string {
	rf := reflect.TypeOf(data)
	var result []string

	for i := 0; i < rf.NumField(); i++ {

		result = append(result, rf.Field(i).Name)
	}
	return result
}

func GetFieldForExcel(data interface{}) []string {
	rf := reflect.TypeOf(data)
	var result []string

	for i := 0; i < rf.NumField(); i++ {

		result = append(result, string(rf.Field(i).Tag.Get("excel")))
	}
	return result
}

// ? extract array 2 dimention to string
func ExtractMultiDimensitinString(data [][]string, callback func([]string)) {
	for i, valueSlice := range data {
		if i == 0 {
			continue
		}
		callback(valueSlice)
	}
}
