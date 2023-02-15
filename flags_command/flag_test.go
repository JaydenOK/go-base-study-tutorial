package flags_command

import (
	"flag"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	env := flag.String("env", "dev", "input run env[dev|test|prod]:")
	flag.Parse()
	fmt.Printf("value:%s \n", *env)
	fmt.Printf("address:%v \n", env)
}
