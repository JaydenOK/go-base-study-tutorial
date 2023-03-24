package ints

import (
	"fmt"
	"testing"
)

//类型说明：
//mysql(1bytes=8bit): TINYINT 1个字节; SMALLINT 2个宇节; MEDIUMINT 3个字节; INT (INTEGHR) 4个字节; BIGINT	8个字节,
//TINYINT	-128〜127	0 〜255
//SMALLINT	-32768〜32767	0〜65535
//MEDIUMINT	-8388608〜8388607	0〜16777215
//INT (INTEGER)	-2147483648〜2147483647	0〜4294967295
//BIGINT	-9223372036854775808〜9223372036854775807	0〜18446744073709551615

//go语言中的int的大小是和操作系统位数相关的，如果是32位操作系统，int类型的大小就是4字节; 如果是64位操作系统，int类型的大小就是8个字节
//go int8、int16、int32 和 int64 即 8、16、32、64 bit，对应mysql

//即 go => mysql 关系
// int8 => TINYINT
// int16 => SMALLINT
// int32 => INT
// int64 => BIGINT

func TestInt(t *testing.T) {
	var a int32
	a = 1234567890
	fmt.Println(a)
}
