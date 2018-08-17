package main


import (
    "fmt"
)



func swap(arr *[]int, x, y int){
    t := (*arr)[x]
    (*arr)[x] = (*arr)[y]
    (*arr)[y] = t
}


func heapify(arr *[]int, length, i int){
    // min heapify
    fmt.Println(arr)
    largest := i
    l := 2 * i + 1
    r := 2 * i + 2
    //fmt.Printf("i: %d, l: %d, r: %d\n", i, l, r)

    if l < length && (*arr)[l] < (*arr)[largest]{
        largest = l
    }

    if r < length && (*arr)[r] < (*arr)[largest]{
        largest = r
    }

    if largest != i {
        swap(arr,i, largest)
        heapify(arr, length, largest)
    }

}


func heapSort(arr *[]int, length int){
    // max sort
    for i := length / 2 - 1; i >= 0 ; i-- {
        fmt.Printf("i: %d , size: %d\n", i, length)
        heapify(arr, length, i)
    }

    for i := length - 1;i >= 0; i-- {
        swap(arr, 0, i)
        heapify(arr, i, 0)
    }

}




func main(){

    var arr = []int{12, 11, 13, 5, 6, 7}

    fmt.Println(arr)
    length := len(arr)

    heapSort(&arr, length)

    fmt.Println(arr)

}
