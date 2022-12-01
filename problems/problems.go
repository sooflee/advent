package problems

import (
	"reflect"
	"strconv"
)

// https://mikespook.com/2012/07/function-call-by-name-in-golang/

var Dailies = map[string]interface{}{
	"Day01": Day01,
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func GetFile(day int) string {
	day_string := strconv.Itoa(day)
	if len(day_string) < 2 {
		day_string = "0" + day_string
	}
	return "day" + day_string + ".in"
}

func GetFunc(day int) string {
	day_string := strconv.Itoa(day)
	if len(day_string) < 2 {
		day_string = "0" + day_string
	}
	return "Day" + day_string
}
