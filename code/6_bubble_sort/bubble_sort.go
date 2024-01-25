// Description: 冒泡排序
package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 3, 2, 5, 4}
	fmt.Println(bubbleSort(arr))
}