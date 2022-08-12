package convert

import (
	"fmt"
	"strconv"
)

func ToBoolP(any interface{}) *bool {
	value := ToBool(any)
	return &value
}

func ToIntP(any interface{}) *int {
	value := ToInt(any)
	return &value
}

func ToInt64P(any interface{}) *int64 {
	value := ToInt64(any)
	return &value
}

func ToFloat32P(any interface{}) *float32 {
	value := float32(ToFloat(any))
	return &value
}

func ToFloat64P(any float64) *float64 {
	value := ToFloat(any)
	return &value
}

func ToStringP(any string) *string {
	value := ToString(any)
	return &value
}

func ToString(any interface{}) (out string) {
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

func ToInt(any interface{}) (out int) {
	return int(ToInt64(any))
}

func ToInt64(any interface{}) (out int64) {
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
		out = ToInt64(0)
	}

	return
}

func ToFloat(any interface{}) (out float64) {
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

func ToBool(any interface{}) (out bool) {
	val := ToString(any)
	out, _ = strconv.ParseBool(val)

	return
}
