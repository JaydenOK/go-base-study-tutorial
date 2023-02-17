package maps

//检查key是否存在
func InMapString(m map[string]string, key string) bool {
	if _, isExist := m[key]; isExist {
		//参数2:true，参数1为其对应值
		return true
	}
	return false
}

//检查key是否存在(整型)
func InMapInt(m map[int]int, key int) bool {
	if _, isExist := m[key]; isExist {
		return true
	}
	return false
}

//合并，键相同，后面的m2覆盖前面m1的值
func Merge(m1 map[string]string, m2 map[string]string) map[string]string {
	for key, value := range m2 {
		m1[key] = value
	}
	return m1
}

// 删除元素，使用内置函数
func Unset(m map[string]string, key string) {
	delete(m, key)
}

// map是无序的，可放入切片排序，输出
func SortNotUseful(m map[string]string) {
}
