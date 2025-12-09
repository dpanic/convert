package convert

// ToSliceOfString converts any slice type to a slice of strings.
// It supports []interface{}, []*string, []string, []int, []int8, []int16,
// []int32, []int64, []float32, []float64, and []bool types.
// Each element is converted using ToString.
// Returns nil for unsupported types.
func ToSliceOfString(any interface{}) (out []string) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []*string:
		for _, item := range any {
			val := ToString(*item)
			out = append(out, val)
		}

	case []string:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []int:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []int8:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []int16:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []int32:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []int64:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []float32:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []float64:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	case []bool:
		for _, item := range any {
			val := ToString(item)
			out = append(out, val)
		}

	default:
	}

	return
}

// ToSliceOfFloat converts any slice type to a slice of float64 values.
// It supports []interface{}, []int, []int8, []int16, []int32, []int64,
// []float32, []float64, and []bool types.
// Each element is converted using ToFloat.
// Returns nil for unsupported types.
func ToSliceOfFloat(any interface{}) (out []float64) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []int:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []int8:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []int16:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []int32:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []int64:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []float32:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []float64:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	case []bool:
		for _, item := range any {
			val := ToFloat(item)
			out = append(out, val)
		}

	default:
		// noop
	}

	return
}

// ToMapOfStrings converts any map type to a map with string keys and string values.
// It supports map[string]interface{} and map[string]string types.
// For map[string]interface{}, each value is converted using ToString.
// Returns an empty map for unsupported types.
func ToMapOfStrings(any interface{}) (out map[string]string) {
	out = make(map[string]string)

	switch val := any.(type) {

	case map[string]interface{}:
		for key, value := range val {
			out[key] = ToString(value)
		}

	case map[string]string:
		out = any.(map[string]string)

	default:
		// noop
	}

	return
}

// ToMapOfInterfaces converts any map type to a map with string keys and interface{} values.
// It supports map[string]interface{} and map[string]string types.
// Returns an empty map for unsupported types.
func ToMapOfInterfaces(any interface{}) (out map[string]interface{}) {
	out = make(map[string]interface{})

	switch val := any.(type) {

	case map[string]interface{}:
		for key, value := range val {
			out[key] = value
		}

	case map[string]string:
		out = any.(map[string]interface{})

	default:
		// fmt.Printf("MapOfStrings, unknown type: %T | %+v\n", val, any)
	}

	return
}

// ToSliceOfMap converts any slice of maps to a slice of map[string]string.
// It supports []map[string]interface{}, []interface{}, and []map[string]string types.
// Each map element is converted using ToMapOfStrings.
// Returns an empty slice for unsupported types.
func ToSliceOfMap(any interface{}) (out []map[string]string) {
	out = make([]map[string]string, 0)

	switch val := any.(type) {

	case []map[string]interface{}:
		for _, item := range val {
			obj := ToMapOfStrings(item)
			out = append(out, obj)
		}

	case []interface{}:
		for _, item := range val {
			obj := ToMapOfStrings(item)
			out = append(out, obj)
		}

	case []map[string]string:
		out = append(out, val...)

	default:
		// noop
	}

	return
}

// ToSliceOfMapOfInterfaces converts any slice of maps to a slice of map[string]interface{}.
// It supports []map[string]interface{}, []interface{}, and []map[string]string types.
// Each map element is converted using ToMapOfInterfaces.
// Returns an empty slice for unsupported types.
func ToSliceOfMapOfInterfaces(any interface{}) (out []map[string]interface{}) {
	out = make([]map[string]interface{}, 0)

	switch val := any.(type) {

	case []map[string]interface{}:
		out = append(out, val...)

	case []interface{}:
		for _, item := range val {
			obj := ToMapOfInterfaces(item)
			out = append(out, obj)
		}

	case []map[string]string:
		for _, item := range val {
			obj := ToMapOfInterfaces(item)
			out = append(out, obj)
		}

	default:
		// noop
	}

	return
}

// ToSliceOfBool converts any slice type to a slice of boolean values.
// It supports []interface{} and []bool types.
// Each element is converted using ToBool.
// Returns nil for unsupported types.
func ToSliceOfBool(any interface{}) (out []bool) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := ToBool(item)
			out = append(out, val)
		}

	case []bool:
		out = append(out, any...)

	default:
		// noop
	}

	return
}
