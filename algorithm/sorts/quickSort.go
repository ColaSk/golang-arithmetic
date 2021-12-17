package sorts

// 递归快排
func qSort(arr []int, low, high int) {

	i := low
	j := high

	if i < j {
		temp := arr[low]
		for i < j {

			for i < j && arr[j] > temp {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}

			for i < j && arr[i] < temp {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}

		}

		arr[j] = temp
		qSort(arr, low, i-1)
		qSort(arr, i+1, high)
	}

}

// 快速排序
func QuickSort(arr []int) {
	i := 0
	j := len(arr) - 1
	qSort(arr, i, j)
}
