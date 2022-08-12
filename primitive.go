package convert

import (
	"fmt"
	"strconv"
)

func BoolP(any interface{}) *bool {
	value := Bool(any)
	return &value
}

func IntP(any interface{}) *int {
	value := Int(any)
	return &value
}

func Int64P(any interface{}) *int64 {
	value := Int64(any)
	return &value
}

func Float32P(any interface{}) *float32 {
	value := float32(Float(any))
	return &value
}

func Float64P(any float64) *float64 {
	value := Float(any)
	return &value
}

func StringP(any string) *string {
	value := String(any)
	return &value
}

func String(any interface{}) (out string) {
	switch any := any.(type) {

	case int:
		out = strconv.Itoa(any)

	case int8:
		out = strconv.FormatInt(int64(any), 10)

	case int16:
		out = strconv.FormatInt(int64(any), 10)

	case int32:
		out = strconv.FormatInt(int64(any), 10)

	case int64:
		out = strconv.FormatInt(any, 10)

	case string:
		out = any

	case float32:
		out = strconv.FormatFloat(float64(any), 'f', -1, 32)

	case float64:
		out = strconv.FormatFloat(any, 'f', -1, 64)

	case bool:
		out = strconv.FormatBool(any)

	default:
		if any == nil {
			out = ""
		} else {
			out = fmt.Sprintf("%v", any)
		}
	}

	return
}

func Int(any interface{}) (out int) {
	return int(Int64(any))
}

func Int64(any interface{}) (out int64) {
	switch any := any.(type) {

	case int:
		out = int64(any)

	case int8:
		out = int64(any)

	case int16:
		out = int64(any)

	case int32:
		out = int64(any)

	case int64:
		out = any

	case float32:
		out = int64(any)

	case float64:
		out = int64(any)

	case string:
		out, _ = strconv.ParseInt(any, 10, 64)

	default:
		out = Int64(0)
	}

	return
}

func Float(any interface{}) (out float64) {
	switch any := any.(type) {

	case int:
		out = float64(any)

	case int8:
		out = float64(any)

	case int16:
		out = float64(any)

	case int32:
		out = float64(any)

	case int64:
		out = float64(any)

	case float32:
		out = float64(any)

	case float64:
		out = any

	case string:
		out, _ = strconv.ParseFloat(any, 64)

	default:
		out = float64(0)
	}

	return
}

func Bool(any interface{}) (out bool) {
	val := String(any)
	out, _ = strconv.ParseBool(val)

	return
}
