package any

import (
	"testing"
)

func TestComposite(t *testing.T) {
	test1 := []int{1, 2, 3, 4, 5}
	res1 := SliceOfFloat(test1)
	if len(res1) == 0 {
		t.Errorf("Result was incorrect, expected length of slice bigger than 0")
	}

	test2 := []int{1, 2, 3, 4, 5}
	res2 := SliceOfString(test2)
	if len(res2) == 0 {
		t.Errorf("Result was incorrect, expected length of slice bigger than 0")
	}

	test8 := map[string]string{
		"key": "val",
	}
	res8 := MapOfStrings(test8)
	if _, ok := res8["key"]; ok == false {
		t.Errorf("Result was incorrect, existance of key 'key'")
	}
}
