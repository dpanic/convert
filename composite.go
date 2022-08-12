package convert

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
		//fmt.Printf("MapOfStrings, unknown type: %T | %+v\n", val, any)
	}

	return
}

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
