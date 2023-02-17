package prints

import (
	"testing"
)

func TestSlice(t *testing.T) {
	type a struct {
		name string
		age  int
	}
	data := a{
		name: "li",
		age:  123,
	}

	//
	var m map[string]interface{}
	m = map[string]interface{}{"a": 1234, "b2": "bbb"}
	PrintVar(data)
	PrintVar(m)
}
