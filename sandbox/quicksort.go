package main

import "fmt"

func main() {
	arr := []int{8, 7, 6, 1, 0, 3, 9, 2, 5, 5, 1, 4, 5}

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low int, high int) int {
	pivot := low - 1

	for low <= high {
		if arr[low] <= arr[high] {
			pivot++

			temp := arr[pivot]
			arr[pivot] = arr[low]
			arr[low] = temp
		}
		low++
	}

	return pivot
}
