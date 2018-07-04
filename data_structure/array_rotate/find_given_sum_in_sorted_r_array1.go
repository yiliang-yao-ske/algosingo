package main

import (
	"fmt"
)

/* 
	Time complexity: O(log n)

 */

func pairInSortedRotated(arr *[]int, length, sum int) bool {
    var i int

    for i = 0; i < length -1; i++{
        if(*arr)[i] > (*arr)[i+1]{
            break
        }
    }
    l := (i+1)%length
    r := i

    for l != r {
        if((*arr)[l] + (*arr)[r] == sum){
            return true
        }
        if((*arr)[l] + (*arr)[r] < sum){
            l = (l + 1) % length
        }else{
            r = (length + r -1) % length
        }
    }
    return false
}


func printArray(a *[]int) { fmt.Println(a) }



func main() {
	var arr1 = []int{11, 15, 26, 38, 1, 9, 10}

	length := len(arr1)
	sum := 16
	printArray(&arr1)

    if(pairInSortedRotated(&arr1, length, sum)){
        fmt.Printf("Array has two elements with sum 16\n")
    }else{
        fmt.Printf("Array doesn't have two elements with sum 16\n")
    }
	return


}
