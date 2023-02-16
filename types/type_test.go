package types

import (
	"fmt"
	"testing"
)

type interger int

var n interger = interger(100)

// var i int = n // cannot use n (type interger) as type int in assignment
var i int = int(n)

func TestT(t *testing.T) {
	fmt.Printf("n = %v, type = %T\n", n, n) //n = 100, type = main.interger
	fmt.Printf("i = %v, type = %T\n", i, i) //i = 100, type = int
}
