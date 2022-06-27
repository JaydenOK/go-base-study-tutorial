package main

import (
	"fmt"
	"time"
)

func GetNowDate() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func main() {
	nowDate := GetNowDate()
	fmt.Println(nowDate)

	timeLayout := "2006-01-02 15:04:05"  //转化所需模板，神奇模板
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, nowDate, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64
	fmt.Println(timestamp)

}
