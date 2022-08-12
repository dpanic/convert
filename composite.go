package any

func SliceOfString(any interface{}) (out []string) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []*string:
		for _, item := range any {
			val := String(*item)
			out = append(out, val)
		}

	case []string:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []int:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []int8:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []int16:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []int32:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []int64:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []float32:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []float64:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	case []bool:
		for _, item := range any {
			val := String(item)
			out = append(out, val)
		}

	default:
	}

	return
}

func SliceOfFloat(any interface{}) (out []float64) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []int:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []int8:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []int16:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []int32:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []int64:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []float32:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []float64:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	case []bool:
		for _, item := range any {
			val := Float(item)
			out = append(out, val)
		}

	default:
		// noop
	}

	return
}

func MapOfStrings(any interface{}) (out map[string]string) {
	out = make(map[string]string)

	switch val := any.(type) {

	case map[string]interface{}:
		for key, value := range val {
			out[key] = String(value)
		}

	case map[string]string:
		out = any.(map[string]string)

	default:
		// noop
	}

	return
}

func MapOfInterfaces(any interface{}) (out map[string]interface{}) {

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

func SliceOfMap(any interface{}) (out []map[string]string) {

	out = make([]map[string]string, 0)

	switch val := any.(type) {

	case []map[string]interface{}:
		for _, item := range val {
			obj := MapOfStrings(item)
			out = append(out, obj)
		}

	case []interface{}:
		for _, item := range val {
			obj := MapOfStrings(item)
			out = append(out, obj)
		}

	case []map[string]string:
		out = append(out, val...)

	default:
		//fmt.Printf("SliceOfMap, unknown type: %T | %+v\n", val, any)
	}

	return
}

func SliceOfMapOfInterfaces(any interface{}) (out []map[string]interface{}) {
	out = make([]map[string]interface{}, 0)

	switch val := any.(type) {

	case []map[string]interface{}:
		out = append(out, val...)

	case []interface{}:
		for _, item := range val {
			obj := MapOfInterfaces(item)
			out = append(out, obj)
		}

	case []map[string]string:
		for _, item := range val {
			obj := MapOfInterfaces(item)
			out = append(out, obj)
		}

	default:
		// noop
	}

	return
}

func SliceOfBool(any interface{}) (out []bool) {
	switch any := any.(type) {

	case []interface{}:
		for _, item := range any {
			val := Bool(item)
			out = append(out, val)
		}

	case []bool:
		out = append(out, any...)

	default:
		// noop
	}

	return
}
