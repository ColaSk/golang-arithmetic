package sorts

// 指定位置插入
func indexInsert(arr []int, i, j, index int) {
	data := arr[j]

	for ; j > index; j-- {
		arr[j] = arr[j-1]
	}
	arr[index] = data

}

// 二分查找位置
func binarySearch(arr []int, i, j int, data int) int {

	if i > j {
		return i
	}

	m := (i + j) / 2

	if arr[m] > data {
		return binarySearch(arr, i, m-1, data)
	} else {
		return binarySearch(arr, m+1, j, data)
	}

}

// 折半插入
func BinaryInsert(arr []int) {

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}
		// 二分查找可能查找到的范围大于查找范围, 说明不需要进行插入
		m := binarySearch(arr, 0, i-1, arr[i])
		if m <= i-1 {
			indexInsert(arr, 0, i, m)
		}

	}
}
