package main

import "fmt"

func linearSeach(data []int, key int) bool {
	for _, item := range data {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{12, 53, 34, 25, 62, 67}
	fmt.Println(linearSeach(items, 24))
}