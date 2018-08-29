package main



import (
    "fmt"

)


func insertSort(arr *[]int, length int) {

    var i, key, j int

    for i = 1; i < length; i++{
        //fmt.Println("update i :", i)
        key = (*arr)[i]
        //fmt.Println("update key : ", key)
        j = i - 1
        //fmt.Println("update j : ", j)

        for j >= 0 && (*arr)[j] > key {
            (*arr)[j + 1] = (*arr)[j]
            //fmt.Println("update arr : ", arr)
            j = j - 1
            //fmt.Println("update j : ", j)
        }
        (*arr)[j + 1] = key
        //fmt.Println("update arr : ", arr)
    }
}




func main(){

    arr := []int{12,11,13,5,6}

    length := len(arr)

    fmt.Println(arr)
    fmt.Println(length)

    insertSort(&arr, length)
    fmt.Println(arr)


}
