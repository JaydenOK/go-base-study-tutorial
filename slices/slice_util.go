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

// 判断字符串是否存在切片
func IntegerInSlice(slice []int, item int) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// 移除存在于切片的字符串
func Remove(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}
