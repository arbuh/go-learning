package main

import "fmt"

func main() {
	arr := []int{8, 7, 6, 1, 0, 9, 2}

	arr = quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) []int {
	return pivot(arr, 0, len(arr)-1)
}

func pivot(arr []int, i int, p int) []int {
	for j := 0; i < len(arr); i++ {
		if arr[i] <= arr[p] {
			temp := arr[j]
			arr[j] = arr[i]
			arr[i] = temp
			j++
		}
	}

	return arr
}
