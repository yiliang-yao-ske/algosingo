package main

import (
	"fmt"
)

/* 
	Time complexity: O(n * d)
	Auxiliary space: O(1)
 */


func leftRotatebyOne(arr *[]int, n int) {
	temp := (*arr)[0]
	var i int

	for i = 0; i < n-1; i++ {
		(*arr)[i] = (*arr)[i+1]
	}
	(*arr)[i] = temp
}

func leftRotate(arr *[]int, d, length int) {

	for i := 0; i < d; i++ {
		leftRotatebyOne(arr, length)

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
