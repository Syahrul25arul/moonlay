package helper

import (
	"reflect"
	"strconv"
	"time"
)

// ? convert value to string
func ConvertFromString(data, typeName string) interface{} {
	if typeName == "int" {
		result, err := strconv.Atoi(data)
		PanicIFError(err)
		return result
	} else if typeName == "uint64" {
		result, err := strconv.ParseUint(data, 10, 64)
		PanicIFError(err)
		return result
	} else if typeName == "time" {
		return ParseTime(data)

	}
	return nil
}

// ? convert data of struct to string
func ConvertToString(data interface{}) []string {

	rf := reflect.TypeOf(data)

	var result []string
	for i := 0; i < rf.NumField(); i++ {
		typeName := reflect.ValueOf(data).Field(i).Type().Name()
		if typeName == "int" {
			data := strconv.Itoa(reflect.ValueOf(data).Field(i).Interface().(int))
			result = append(result, data)

		} else if typeName == "int64" {
			data := strconv.FormatInt(reflect.ValueOf(data).Field(i).Interface().(int64), 10)
			result = append(result, data)
		} else if typeName == "Time" {
			time := reflect.ValueOf(data).Field(i).Interface().(time.Time)
			data := ParseTimeToString(time)
			result = append(result, data)
		} else {
			result = append(result, reflect.ValueOf(data).Field(i).String())
		}

	}
	return result
}
