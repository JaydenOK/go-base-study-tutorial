package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	str := "test-Testet"
	//a := word[0:12]
	result := strings.HasPrefix(str, "te")
	fmt.Println(result)
}

// 字符串方法介绍
func usefulFunctions() {
	str := "test-Test"                   //go语言字符串的本质就是byte[]数组，里面每一个数据存的是字符的Unicode码，可以用读取字符串字符形式读取
	_ = str[1]                           //获取指定位置的字符
	_ = len(str)                         //获取字符串长度
	_ = str[1:2]                         //截取子字符串，[start:length], start从0开始,length 不能超过字符串最长长度,可先用len判断长度
	strings.Contains(str, "st")          //字符串包含判断；字符串是否包含在内，区分大小写
	strings.ContainsAny(str, "aT")       //字符串任意包含判断；判断字符串s中是否包含个子串str中的"任何一个"字符。包含则返回true
	strings.Index(str, "e")              //获取字符下标；判断第一个出现的位置，从0开始，不存在 返回-1
	strings.Split("a-b-c-d-e", "-")      //字符串分割，与Join相反
	strings.Join([]string{str, str}, "") //字符串切片转字符串；用空把数组拼接成字符串
	strings.Replace(str, "t", "T", -1)   //字符串替换；最后参数表示起始位置和长度，-1表示到结尾
	strings.ToLower(str)                 //字符串转小写
	strings.ToUpper(str)                 //字符串转大写
	strings.TrimRight(str, "et")         //删除后端包含cutset字符的东西，直到没有cutset字符匹配为止
	strings.HasPrefix(str, "et")         //判断是否以某个字符串为前缀

	strings.Count(str, "t")      //函数第二个参数中指定字符串的个数
	strings.HasPrefix(str, "te") //判断前缀
	strings.HasSuffix(str, "st") //判断后缀
	strings.LastIndex(str, "t")  //判断最后一个出现的位置
	strings.Repeat("a", 5)       //字符串重复5次

}

// Go 语言中的字符串是一个字节切片。可以像切片一样使用，下标访问修改，Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。
func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x-", s[i]) //	%x输出十六进制编码
		//fmt.Printf("%c-", s[i])  	// %c输出字符
	}
}

// rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
func printChars(s string) {
	//传入字符串，字符串切片 转成 rune切片
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c-", runes[i])
	}
}

func bade() {
	name := "Hello World"
	printBytes(name) //48-65-6c-6c-6f-20-57-6f-72-6c-64-
	fmt.Printf("\n")

	printChars(name) //H-e-l-l-o- -W-o-r-l-d-
	fmt.Printf("\n\n")
	//这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码占用了 c3 和 b1 两个字节。它的 UTF-8 编码占用了两个字节 c3 和 b1。而我们打印字符时，
	//却假定每个字符的编码只会占用一个字节，这是错误的。在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间

	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")

	printChars(name)
	fmt.Printf("\n")

	//for range 循环字符串，循环返回的是是当前 rune 的字节位置
	for index, rune := range name {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}

	runeToString()
	editString()
}

// 十六进制字节切片 转 字符串，直接十进制转也可以
func toString() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
}

// rune双字节切片 转 字符串
func runeToString() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println(str)                         //Señor
	fmt.Println(len(str))                    //6
	fmt.Println(utf8.RuneCountInString(str)) //5
}

// go 字符串 修改方法(Go 中的字符串是不可变的)
func editString() {
	h := "hello"
	//不能直接修改字符串  h[0] = 'a' (×)
	runeString := []rune(h)
	fmt.Println("runeSring:", runeString)
	runeString[0] = 'a'
	h = string(runeString)
	fmt.Println("h:", h)
}

func demo() {
	//一、字符串的操作方法
	//1、len(str)  --- 统计字符串长度
	//2、r := []rune(str) --- 字符串遍历同时处理有中文的问题
	//3、n, err := strconv.Atoi(str)   ---- 字符串转整数
	//4、str = strconv.Itoa(被转后的数字) --- 整数转字符串
	//5、var bytes = []byte(str)   --- 字符串转[]byte
	//6、stringstrings.Contains("字符串one", "字符串two") --- 查找字符串two是否在字符串one中返回bool值
	//7、stringstrings.count("字符串one", "字符串two")  ----  统计字符串two在字符串one中出现的次数

	//1.1统计字符串的长度 len()
	//统计字符串的长度，按字节len(str)
	//golang的编码统一都是utf-8(ascii的字符（字母和数字）占一个字节，汉字占用三个字节)
	//strOne := "老王"
	//strtwo := "laowang"
	//fmt.Println(len(strOne)) // 6
	//fmt.Println(len(strtwo)) // 7
	//
	//
	////1.2、字符串遍历，同时处理有中文的问题
	////大家好奇为什么要这样呢？直接类似python和java一样循环打印不就行了吗，答案是go不支持这样直接遍历，需要转换
	////Golang中所有的字符都是以utf-8编码存储的，对于中文字符来说，一个中文字符占3个字节。用传统方式输出的话会出现中文乱码，原因是传统方式是以字节的方式进行遍历的，而中文字符占了3个字节
	////例如直接遍历字符串
	//strtest := "laowang 你是一个好人"
	//for i := 0; i < len(strtest); i++ {
	//	fmt.Printf("%c", strtest[i])  // 输出的是乱码
	//
	//}
	//
	////1.2.1、遍历带有中文的字符串方式一： str = []rune(str)
	////字符串遍历，同时处理有中文的问题 r := []rune(str)
	//str2 := "hello北京"
	//r := []rune(str2)
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("字符=%c\n", r[i])
	//}
	//
	////1.2.2、方式二使用for…range
	//strtest := "laowang 你是一个好人"
	//for index, val := range strtest {
	//	fmt.Printf("%d, %c", index, val)
	//}
	//
	////1.3、字符串转整数：strconv.Atoi(str)
	//strconv.Atoi(str), //传递进来一个只有数字的字符串，返回的是int类型和err信息
	//// 字符串转整数
	//// strconv.Atoi(str), 传递进来一个只有数字的字符串，返回的是int类型和err信息
	//n, err := strconv.Atoi("123")  // n就是转化后的int数值
	//if err != nil { // nil可以暂时理解为null
	//	fmt.Println("转换有误", err)
	//} else {
	//	fmt.Println("转换的成果是: ", n)
	//}
	//
	//1.4、整数转字符串 :strconv.Itoa(int)
	//strconv.Itoa(int), 传递一个int类型返回str
	//strnum := strconv.Itoa(123)
	//fmt.Println("转化后的字符串", strnum) // 转化后的字符串 123
	//fmt.Printf("类型是: %T", strnum) //string
	//
	//
	//1.5、字符串 转 []byte: var bytes = []byte(str)
	//var bytes = []byte("hello go")
	//fmt.Printf("bytes=%v\n", bytes)  // 输出的是ascii码
	//
	//
	//1.6、[]byte 转 字符串: str = string([]byte{byte, byte, byte})
	//str := string([]byte{97, 98, 99})
	//fmt.Println(str) // abc
	//
	//
	//1.7、//查找子串是否在指定的字符串中: stringstrings.Contains()
	//stringstrings.Contains(“被查的总字符串”,“要查的字符串”)
	//fmt.Println(stringstrings.Contains("laowang123", "laow"))  // true
	//
	//
	//1.8、统计一个字符串有几个指定的子串 ： stringstrings.Count(“str”, “substr”)
	//stringstrings.Count(“sunstr”, “str”) 返回的是int
	//fmt.Println(stringstrings.Count("老王", "王"))     // 1
	//fmt.Println(stringstrings.Count("lwaong", "王")) // 0
	//fmt.Println(stringstrings.Count("老王头王", "王")) // 2
	//
	//
	//1.9、stringstrings.EqualFold（str1,str2），不区分大小写的字符串比较(==是区分字母大小写的)
	//strnumone := "laowang"
	//strnumtwo := "Laowang"
	//fmt.Println(strnumone == strnumtwo) // false
	//fmt.Println(stringstrings.EqualFold(strnumone, strnumtwo))  // true
	//
	//
	//1.10、返回子串在字符串第一次出现的index值，如果没有返回-1 :
	//fmt.Println(stringstrings.Index("我是老王头", "王"))      // 9, 因为中文占三个字节所以第十位下标是9
	//fmt.Println(stringstrings.Index("laowangtou", "w")) // 3
	//fmt.Println(stringstrings.Index("老王头", "你")) // -1
	//
	//
	//1.11、stringstrings.LastIndex（）返回子串在字符串最后一次出现的index，
	//fmt.Println(stringstrings.LastIndex("老王头", "王"))     // 3
	//fmt.Println(stringstrings.LastIndex("laowang", "a")) // 4
	//fmt.Println(stringstrings.LastIndex("老王是你", "不"))  // -1
	//
	//
	//
	//1.12、将指定的子串替换成 另外一个子串: stringstrings.Replace（）
	//将指定的子串替换成 另外一个子串: stringstrings.Replace(总字符串, 替换的值，替换后的值，替换的次数）
	//stringstrings.Replace(“字符串”, ''要替换的字符", “替换后的值” -1 )-1 代表全部替换
	//fmt.Println(stringstrings.Replace("gogo hello", "go", "go语言", 2)) // go语言go语言 hello ， 将gogohello这个字符串的go都替换成go语言，替换两次
	//
	//fmt.Println(stringstrings.Replace("gogogogo", "go", "go语言", -1))  // -1代表全部替换
	//
	//
	//
	//1.13、字符串转化为数组 string.Split()
	//和python一样使用split切割，将字符串按照某个字符切割为数组
	//格式：
	//stringstrings.Split(str, “按照某个字符切割”)
	//strSeptember := "老王,你可,少点,套路吧"
	//var arr []string = stringstrings.Split(strSeptember, ",")
	//fmt.Println(arr) // []string
	//fmt.Printf("%T\n", arr)  // [老王 你可 少点 套路吧]
	//
	//
	////stringstrings.Split("hello,wrold,ok", ",")
	//strArr := stringstrings.Split("hello,wrold,ok", ",")
	//for i := 0; i < len(strArr); i++ {
	//	fmt.Printf("str[%v]=%v\n", i, strArr[i])
	//}
	//fmt.Printf("strArr=%v\n", strArr)
	//
	//
	//1.14、将字符串的字母进行大小写的转换:
	//stringstrings.ToUpper (str) # 将字符串中的小写字母全部大写
	//stringstrings.ToLower (str) # 将字符串中的大写字母全部小写
	//strUpper := "laowang是一个大流氓"
	//strUpperto := stringstrings.ToUpper(strUpper) //全部大写, 是有返回值的 必须用一个str元素接收
	//fmt.Println(strUpperto)                 // LAOWANG是一个大流氓
	//
	//strLower := "LAOWANG住在隔壁"
	//strLowerto := stringstrings.ToLower(strLower)  // 全部小写
	//fmt.Println(strLowerto) // laowang住在隔壁
	//
	//
	//1.15、将字符串左右两边的空格去掉： stringstrings.TrimSpace(str)
	//stringstrings.TrimSpace（str）
	//strTrim := " kongge "
	//fmt.Println(stringstrings.TrimSpace(strTrim))
	//
	//
	//1.16、将字符串左右两边指定的字符去掉
	//stringstrings.Trim(“str”, “要去除的字符”)
	//strTrim := "alaowanga"
	//fmt.Println(stringstrings.Trim(strTrim, "a")) //laowang
	//
	////1.17、只去除左边的指定字符
	//stringstrings.TrimLeft(str, cutsetstr)
	//strTrim := "alaowanga"
	//fmt.Println(stringstrings.TrimLeft(strTrim, "a")) //laowanga ，  去除左边的a字符
	//
	////1.18、只去除右边的指定字符
	//stringstrings.TrimRight(str,cutset )
	//strTrim := "alaowanga"
	//fmt.Println(stringstrings.TrimRight(strTrim, "a")) //alaowang 去除右边的a字符
	//
	//
	////1.19、查询字符所在的下标 index
	////查不到返回-1
	//strIndex := "老王"
	//fmt.Println(stringstrings.IndexAny(strIndex, "王")) //3
	//fmt.Println(stringstrings.Index(strIndex, "王"))  // 3
	//strIndexA := "俺是Laowang老王头"
	//fmt.Println(stringstrings.Index(strIndexA, "h")) //  -1
	//
	//
	////1.17、判断字符串是否以指定的字符串开头: stringstrings.HasPrefix
	//stringstrings.HasPrefix(str,prefix) //返回bool值
	//
	//strIndexA := "俺是Laowang老王头"
	//
	//fmt.Println(stringstrings.HasPrefix(strIndexA, "a"))
	//fmt.Println(stringstrings.HasPrefix(strIndexA, "俺"))
}
