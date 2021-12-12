package sorts

// 希尔排序(缩小增量排序)
func ShellSort(arr []int) []int {

	length := len(arr)
	gap := length / 2

	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		// 缩小增量
		gap = gap / 2
	}
	return arr
}
