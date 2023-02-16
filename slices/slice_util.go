package slices

// 判断字符串是否存在切片
func StringInSlice(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 判断数字是否存在切片
func IntegerInSlice(slice []int, item int) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 切片截取，从下标index开始，截取长度length
func Cut(slice []string, index int, length int) []string {
	//len(slice)当前容量, cap(slice)总容量，包括空值nil
	maxLength := len(slice)
	endLength := index + length
	if endLength > maxLength {
		endLength = maxLength
	}
	return slice[index:endLength]
}

// 字符串切片合并
func Merge(slice []string, slice2 []string) []string {
	slice3 := append(slice, slice2)
	return slice3
}

// 移除存在于切片的字符串
func Remove(slice []string, item string) bool {
	_ = make([]string, len(slice))
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
