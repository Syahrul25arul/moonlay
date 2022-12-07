package helper

import (
	"moonlay/model/domain"
	"reflect"
	"strconv"
	"time"
)

func GetField(data interface{}) []string {
	rf := reflect.TypeOf(data)
	var result []string

	for i := 0; i < rf.NumField(); i++ {

		result = append(result, rf.Field(i).Name)
	}
	return result
}

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

func ChooseDomain(nameDomain string) *domain.Datamart1 {
	if nameDomain == "Datamart1" {
		return &domain.Datamart1{}
	}
	return nil
}
