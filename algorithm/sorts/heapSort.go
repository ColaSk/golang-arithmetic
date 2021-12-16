package sorts

/*完成数组啊让人a[low]到a[high]的范围内对在位置low上的节点进行调整*/
func sift(arr *[]int, low, high int) {
	i := low
	j := 2 * i
	temp := (*arr)[i]
	for j <= high {
		if j < high && (*arr)[j] < (*arr)[j+1] {
			j++
		}

		if temp < (*arr)[j] {
			(*arr)[i] = (*arr)[j]
			i = j
			j = 2 * i
		} else {
			break
		}
	}
	(*arr)[i] = temp

}

// 堆排序
// 支持 1---n-1的排序
func HeapSort(arr []int) []int {
	l := len(arr) - 1

	for i := l / 2; i >= 1; i-- {
		sift(&arr, i, l)
	}

	for i := l; i >= 2; i-- {
		arr[1], arr[i] = arr[i], arr[1]
		sift(&arr, 1, i-1)
	}

	return arr
}
