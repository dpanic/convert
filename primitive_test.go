package any

import (
	"testing"
)

func TestPrimitive(t *testing.T) {
	test3 := "true"
	res3 := Bool(test3)
	if res3 == false {
		t.Errorf("Result was incorrect, expected true")
	}

	test4 := 1
	res4 := Float(test4)
	if res4 == 0 {
		t.Errorf("Result was incorrect, expected > 0")
	}

	test5 := 1
	res5 := Int(test5)
	if res5 == 0 {
		t.Errorf("Result was incorrect, expected > 0")
	}

	test6 := 1
	res6 := Int(test6)
	if res6 == 0 {
		t.Errorf("Result was incorrect, expected > 0")
	}

	test7 := "string"
	res7 := String(test7)
	if res7 != "string" {
		t.Errorf("Result was incorrect, expected 'string' value, got '%s'", res7)
	}
}
