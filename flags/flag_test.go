package flags

import (
	"flag"
	"fmt"
	"testing"
)

func Test_assert(t *testing.T) {
	flagConf := flag.String("conf", "conf/wmqx.toml", "please input conf path")
	fmt.Println(flagConf)
}
