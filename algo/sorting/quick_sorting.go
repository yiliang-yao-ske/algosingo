package main



import (
    "fmt"

)

func swap(arr *[]int, a, b int)  {
    (*arr)[a], (*arr)[b] = (*arr)[b], (*arr)[a]
}



func partition(arr *[]int, start, end int) (p int) {

    x := (*arr)[end]
    i := start - 1

    for j := start; j <= end - 1; j++ {
        if (*arr)[j] <= x {
            i++
            swap(arr, i, j)
        }
    }
    swap(arr, i+1, end)
    return i + 1
}




func quickSort(arr *[]int, start, end int)  {

    if start < end {
        p := partition(arr, start, end)
        fmt.Println("p: ", p)
        fmt.Println(start, p-start)
        quickSort(arr, start, p - 1)
        quickSort(arr, p + 1, end)
    }

}



func main(){

    arr := []int{12,11,13,5,6}

    length := len(arr)

    fmt.Println(arr)
    fmt.Println(length)

    quickSort(&arr, 0, length -1)
    fmt.Println(arr)


}
