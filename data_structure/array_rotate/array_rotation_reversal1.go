package main

import (
	"fmt"
)

/* 
	Time complexity: O(n)
	Auxiliary space: O(1)
 */


func leftRotate(arr *[]int, d, length int) {

	reverseArray(arr, 0, d-1)
	reverseArray(arr, d, length-1)
	reverseArray(arr, 0, length-1)
}

func reverseArray(arr *[]int, start, end int) {
	var temp int

	for start < end {
		temp = (*arr)[start]
		(*arr)[start] = (*arr)[end]
		(*arr)[end] = temp
		start++
		end--

	}
}

func printArray(a *[]int) { fmt.Println(a) }

func main() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7}

	length := len(arr1)

	printArray(&arr1)

	leftRotate(&arr1, 2, length)

	printArray(&arr1)

}
