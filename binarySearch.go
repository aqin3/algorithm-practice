package main

import "fmt"

func main() {
	data := []int{12, 23, 34, 45, 56, 67, 78, 89, 90}
	fmt.Println(binarySearch(data, 12))
}

// function returns index of the target
func binarySearch(data []int, target int) int {
	low, high := 0, len(data)-1

	if target == data[high] {
		return high
	}

	for high > low {
		mid := (high + low) / 2
		if target == data[mid] {
			return mid
		}

		if target < data[mid] {
			high = mid
		} else {
			low = mid
		}
	}

	return -1
}
