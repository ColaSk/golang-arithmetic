package sorts

// 插入操作
func insert(arr *[]int, i, j int, data int) {
	// 此处对迭代i，j进行操作，不太合理可能造成理解混乱
	// TODO: 优化实现方式
	for ; i < j; i++ {
		// 移位 并插入
		if (*arr)[i] > data {
			for ; j > i; j-- {
				(*arr)[j] = (*arr)[j-1] // 数组指针取值赋值操作
			}
			(*arr)[i] = data
		}
	}
}

// 插入排序
func InsertSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		insert(&arr, 0, i, arr[i])
	}
	return arr
}
