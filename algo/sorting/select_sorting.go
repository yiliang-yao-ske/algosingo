package main



import (
    "fmt"
    //"math"
)


func selectSort(arr *[]int, length int) {

    var i int

    var key int = 0

    var left int = 0

    for left = 0; left < length; left++{


        for i = left; i < length; i++{
            key = 0 + left
            // find minimum element in array
            if (*arr)[i] < (*arr)[key] {
                key = i
            }
            //fmt.Println("update arr : ", arr)
        }
        fmt.Println("key: ",key)
        fmt.Println("arr[key]: ",(*arr)[key])
        swap(arr, left, key)

    }
}

func swap(arr *[]int, left, right int){

    (*arr)[left], (*arr)[right] =  (*arr)[right], (*arr)[left]

}


func main(){

    arr := []int{64, 25, 12, 22, 11}

    length := len(arr)

    fmt.Println(arr)
    fmt.Println(length)

    selectSort(&arr, length)
    fmt.Println(arr)


}
