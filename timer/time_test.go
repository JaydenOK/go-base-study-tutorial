package timer

import (
	"fmt"
	"testing"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"
const DateFormat = "2006-01-02"

func Test_time(t *testing.T) {
	fmt.Println(GetCurrentTimestamp())      //11位
	fmt.Println(GetCurrentMilliTimestamp()) //14位
	fmt.Println(GetCurrentMicroTimestamp()) //17位
	fmt.Println(GetCurrentDateTime())
	fmt.Println(GetCurrentDate())
}

// 获取当前时间戳 1664531446
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// 获取毫秒时间戳 1664531446277
func GetCurrentMilliTimestamp() int64 {
	return time.Now().UnixMilli()
}

// 获取毫秒时间戳 1664531446277667
func GetCurrentMicroTimestamp() int64 {
	return time.Now().UnixMicro()
}

// 获取当前日期时间 2022-09-30 17:50:46
func GetCurrentDateTime() string {
	return time.Now().Format(TimeFormat)
}

// 获取当前日期 2022-09-30
func GetCurrentDate() string {
	return time.Now().Format(DateFormat)
}
