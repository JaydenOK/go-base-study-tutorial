package tests

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcde"
	fmt.Println(string(s[2]))
}
