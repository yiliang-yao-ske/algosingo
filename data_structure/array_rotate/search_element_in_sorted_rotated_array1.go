package main

import (
	"fmt"
)

/* 
	Time complexity: O(log n)

 */



func leftRotate(arr *[]int, d, length int) {
	var i, j, k, temp int

	for i = 0; i < gcd(d, length); i++ {
		temp = (*arr)[i]
		j = i
		for true{
			k = j + d
			fmt.Printf("k: %d\n", k)
			if (k >= length) {k = k -length}
			if (k == i) {break}
			(*arr)[j] = (*arr)[k]
			j = k
		}
		(*arr)[j] = temp
	}

}

func printArray(a *[]int) { fmt.Println(a) }

func gcd(x, y int) int{
	for y != 0 {
		x, y = y, x %y
	}
	return x
}


func main() {
	var arr1 = []int{5, 6, 7, 8, 9, 10, 1, 2, 3}

	length := len(arr1)

	printArray(&arr1)

	leftRotate(&arr1, 2, length)

	printArray(&arr1)

}
