# Convert ANY to ANY type in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/dpanic/conv)](https://goreportcard.com/report/github.com/dpanic/any)

Converting String to []string, or string to float64 could be handy sometimes, when your project has inconsistent input.

## Features
* Convert basic types to any other basic type
* Same as above but with * pointer

### Primitive type conversions:
* BoolP -> *bool
* IntP -> *int
* Int64P -> *int64
* Float32P -> *float32
* Float64P -> *float64
* StringP -> *string
* String -> string
* Int -> int
* Int64 -> int64
* Float -> float64
* Bool -> bool

### Composite type conversions:
* SliceOfString -> []string
* SliceOfFloat -> []float64
* MapOfStrings -> map[string]string
* MapOfInterfaces -> map[string]interface{}
* SliceOfMap -> []map[string]string
* SliceOfMapOfInterfaces -> []map[string]interface{}
* SliceOfBool -> []bool
