package search

// 折半查找

func BinarySearch(arr []int, target int, low, high int) (int, bool) {
	if low > high || len(arr) == 0 {
		return -1, false
	}

	mid := (low + high) / 2

	if target > arr[mid] {
		return BinarySearch(arr, target, mid+1, high)
	} else if target < arr[mid] {
		return BinarySearch(arr, target, low, mid-1)
	} else {
		return mid, true
	}
}
