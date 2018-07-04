package main

import (
	"fmt"
)

/* 
	Time complexity: O(log n)

 */

func pairInSortedRotated(arr *[]int, length, sum int) int {
    var i int

    for i = 0; i < length -1; i++{
        if(*arr)[i] > (*arr)[i+1]{
            break
        }
    }
    l := (i+1)%length
    r := i

    count := 0

    for l != r {
        if((*arr)[l] + (*arr)[r] == sum){
            count++

            if(l == (r -1 + length)%length){
                return count
            }
            l = (l + 1) %length
            r = (r - 1 + length)%length

        }else if((*arr)[l] + (*arr)[r] < sum){
            l = (l + 1) % length
        }else{
            r = (length + r -1) % length
        }
    }
    return count
}


func printArray(a *[]int) { fmt.Println(a) }



func main() {
	var arr1 = []int{11, 15, 26, 38, 1,7 ,9, 10}

	length := len(arr1)
	sum := 16
	printArray(&arr1)

    fmt.Printf("%d\n", pairInSortedRotated(&arr1, length, sum))
	return


}
