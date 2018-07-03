package main

import (
	"fmt"
)

/* 
	Time complexity: O(log n)

 */

func pivotedBinarySearch(arr *[]int, n, key int) int{
	pivot := findPivot(arr, 0, n - 1)

	if(pivot == -1){
		return binarySearch(arr, 0, n-1, key)
	}

	if((*arr)[pivot] == key){return pivot}
	if((*arr)[0] <= key){
		return binarySearch(arr, 0, pivot -1, key)
	}
	return binarySearch(arr, pivot + 1, n -1, key)
}


func findPivot(arr *[]int, low, high int) int {
	if(high < low){return -1}
	if(high == low){return low}

	mid := (low + high)/2
	if(mid < high && (*arr)[mid] > (*arr)[mid + 1]){
		return mid
	}
	if(mid > low && (*arr)[mid] < (*arr)[mid -1]){
		return (mid -1)
	}
	if((*arr)[low] >= (*arr)[mid]){
		return findPivot(arr, low, mid-1)
	}
	return findPivot(arr, mid + 1, high)
}

func printArray(a *[]int) { fmt.Println(a) }

func binarySearch(arr *[]int, low, high, key int) int{
	if(high < low){ return -1}
	mid := (low + high) / 2
	if(key == (*arr)[mid]){
		return mid
	}
	if (key > (*arr)[mid]){
		return binarySearch(arr, (mid + 1), high, key)
	}
	return binarySearch(arr, low, (mid - 1), key)
}


func main() {
	var arr1 = []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}

	length := len(arr1)
	key := 9
	printArray(&arr1)

	fmt.Printf("Index of the element is : %d\n",
	pivotedBinarySearch(&arr1, length, key))
	return


}
