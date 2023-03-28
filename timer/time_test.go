package timer

import (
	"app/utils"
	"fmt"
	"testing"
)

func Test_time(t *testing.T) {

	var s string
	s = "2021-03-20 01:21:27"
	result := utils.ParseDateToTime(s)
	fmt.Println(result)

	result2 := utils.FormatTimeToDate(int(result))
	fmt.Println(result2)
}
