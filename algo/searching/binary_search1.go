package main

import (
	"fmt"
)

func binarySearch(arr *[]int, left, right, x int) int{
	if ( right >= left){
		mid := left + ( right -1) /2

		if((*arr)[mid] == x){
			return mid
		}
		if((*arr)[mid] > x){
			return binarySearch(arr, left, mid-1, x)
		}
		return binarySearch(arr, mid+1,  right, x)
	}
	return -1
}

func main() {
	var arr1 = []int{2,3,4,10,40}
	length := len(arr1)
	var x int = 10
	fmt.Println(arr1)
	result := binarySearch(&arr1, 0, length -1 , x)
	if(result == -1){
		fmt.Printf("Element is not present in array\n")
	}else{
		fmt.Printf("Element is presesnt at index %d\n", result)
	}

}