// Package convert provides type conversion utilities for Go primitive and composite types.
// It offers safe type conversions between various Go types including integers, floats,
// strings, booleans, slices, and maps. All conversion functions handle nil values
// gracefully and return zero values when conversion is not possible.
package convert

import (
	"fmt"
	"strconv"
)

// ToBoolP converts any value to a boolean and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToBoolP(any interface{}) *bool {
	value := ToBool(any)
	return &value
}

// ToIntP converts any value to an int and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToIntP(any interface{}) *int {
	value := ToInt(any)
	return &value
}

// ToInt64P converts any value to an int64 and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToInt64P(any interface{}) *int64 {
	value := ToInt64(any)
	return &value
}

// ToFloat32P converts any value to a float32 and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToFloat32P(any interface{}) *float32 {
	value := float32(ToFloat(any))
	return &value
}

// ToFloat64P converts a float64 value and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToFloat64P(any float64) *float64 {
	value := ToFloat(any)
	return &value
}

// ToStringP converts a string value and returns a pointer to the result.
// It is useful when working with APIs that require pointer types.
func ToStringP(any string) *string {
	value := ToString(any)
	return &value
}

// ToString converts any value to its string representation.
// It supports int, int8, int16, int32, int64, string, float32, float64, and bool types.
// For unsupported types, it uses fmt.Sprintf with %v format.
// Returns an empty string for nil values.
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

// ToInt converts any value to an int.
// It delegates to ToInt64 and casts the result to int.
// Returns 0 for unsupported types or nil values.
func ToInt(any interface{}) (out int) {
	return int(ToInt64(any))
}

// ToInt64 converts any value to an int64.
// It supports int, int8, int16, int32, int64, float32, float64, and string types.
// For string values, it attempts to parse as base-10 integer.
// Returns 0 for unsupported types, nil values, or invalid string formats.
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

// ToFloat converts any value to a float64.
// It supports int, int8, int16, int32, int64, float32, float64, and string types.
// For string values, it attempts to parse as a 64-bit floating point number.
// Returns 0 for unsupported types, nil values, or invalid string formats.
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

// ToBool converts any value to a boolean.
// It first converts the value to a string using ToString, then parses it as a boolean.
// Accepted string values are: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Returns false for unsupported values or invalid string formats.
func ToBool(any interface{}) (out bool) {
	val := ToString(any)
	out, _ = strconv.ParseBool(val)

	return
}
