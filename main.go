package main

import (
	"arithmetic/algorithm/sorts"
	"log"
)

func main() {
	log.Printf("zzz")
	var test_list = []int{2, 1, 3, 45, 78, 40, 2}
	sorts.QuickSort(test_list)
	log.Print(test_list)
}
