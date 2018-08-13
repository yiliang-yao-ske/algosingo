package main


import (
    "fmt"
)



type Heap struct{
    h_arr *[]int
    capacity int
    h_size int
}

func (heap *Heap)parent(i int) int{
    return (i - 1)/2
}

func (heap *Heap)left(i int) int{
    return (2 * i + 1)
}

func (heap *Heap)right(i int) int{
    return (2 * i + 2)
}

func (heap *Heap)getMin() int{
    return (*heap.h_arr)[0]
}


func (heap *Heap)insertKey(k int){
    if heap.h_size == heap.capacity {
        fmt.Println("Overflow: Could not insertKey")
        return
    }

    heap.h_size += 1
    i := heap.h_size -1
    (*heap.h_arr)[i] = k

    // fmt.Printf("parent value : %d\n", (*heap.h_arr)[heap.parent(i)])
    // fmt.Printf("inserted value : %d\n", (*heap.h_arr)[i])
    // fmt.Printf("%t \n", (*heap.h_arr)[heap.parent(i)] > (*heap.h_arr)[i])
    for i != 0 && (*heap.h_arr)[heap.parent(i)] > (*heap.h_arr)[i] {
        swap((*heap.h_arr)[i], (*heap.h_arr)[heap.parent(i)])
        i = heap.parent(i)
    }

}


func (h *Heap)sort(i int){
    // max sort


}




func swap(x, y int)(int, int){
    return y, x
}


func createHeap(cap int)(*Heap){
    heap_array := make([]int, cap)
    return&Heap{
        h_size : 0,
        capacity : cap,
        h_arr : &heap_array,
    }
}


func main(){
    heap := createHeap(11)

    heap.insertKey(3)
    fmt.Println(heap.h_arr)
    heap.insertKey(2)
    fmt.Println(heap.h_arr)
    heap.deleteKey(1)
    fmt.Println(heap.h_arr)
    heap.insertKey(15)
    fmt.Println(heap.h_arr)
    heap.insertKey(5)
    fmt.Println(heap.h_arr)
    heap.insertKey(4)
    fmt.Println(heap.h_arr)
    heap.insertKey(45)
    fmt.Println(heap.h_arr)
    heap.minheapify(2)
    fmt.Println(heap.h_arr)
    fmt.Println(heap.extractMin())
    fmt.Println(heap.h_arr)
    fmt.Println(heap.getMin())
    fmt.Println(heap.h_arr)
    heap.decreaseKey(2, 1)
    fmt.Println(heap.getMin())

}
