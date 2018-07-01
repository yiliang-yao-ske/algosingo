package main

import (
	"fmt"
)

/* 
	This solution has problem that:
		swap base should change during rotation
	Time complexity: O(n)
	Auxiliary space: O(1)
 */


func leftRotate(arr *[]int, d, length int) {
	if(d == 0 || d == length) {return}
	if(length - d == d){
		swap(arr, 0, length-d, d)
		return
	}

	if(d < length-d){
		swap(arr, 0, length-d, d)
		leftRotate(arr, d, length-d)
	}else{
		swap(arr, 0, d, length-d)
		leftRotate(arr, length-d, d)
	}
}

func swap(arr *[]int, fi, si, d int) {
	var i, temp int
	for i = 0; i < d; i++{
		temp = (*arr)[fi + i]
		(*arr)[fi + i] = (*arr)[si + i]
		(*arr)[si + i] = temp
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
