package sorts

// 归并操作
func merge(arr []int, low, mid, high int) {
	i := mid
	j := mid + 1

	for j <= high {

		// 插入排序
		temp := arr[j]
		for i >= low && temp < arr[i] {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = temp
		i = j
		j++

	}
}

// 二路归并排序
func MergeSort(arr []int, low, high int) {

	mid := (low + high) / 2

	if low < high {
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}

}
