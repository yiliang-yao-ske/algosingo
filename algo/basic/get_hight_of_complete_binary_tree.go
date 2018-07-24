package main

import (
	"fmt"
	"math"
)

func get_height(n int) int{

	return int(math.Ceil(math.Log2(float64(n + 1))) - 1)
}


func main() {
	// N is number of node number of certain complete binary tree
	N := 6
	fmt.Println(get_height(N))
}