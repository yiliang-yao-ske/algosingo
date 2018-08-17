package main


import (
    "fmt"
)



type Heap struct{
    arr *[]int
    capacity int
    size int
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
    return (*heap.arr)[0]
}


func (heap *Heap)insertKey(k int){
    if heap.size == heap.capacity {
        fmt.Println("Overflow: Could not insertKey")
        return
    }

    heap.size += 1
    i := heap.size -1
    (*heap.arr)[i] = k

    // fmt.Printf("parent value : %d\n", (*heap.arr)[heap.parent(i)])
    // fmt.Printf("inserted value : %d\n", (*heap.arr)[i])
    // fmt.Printf("%t \n", (*heap.arr)[heap.parent(i)] > (*heap.arr)[i])
    // for i != 0 && (*heap.arr)[heap.parent(i)] > (*heap.arr)[i] {
    //     swap((*heap.arr)[i], (*heap.arr)[heap.parent(i)])
    //     i = heap.parent(i)
    // }

    heap.heapify(heap.size, i)

}


func swap(arr *[]int, x, y int){
    t := (*arr)[x]
    (*arr)[x] = (*arr)[y]
    (*arr)[y] = t
}


func (h *Heap)heapify(length, i int){
    // min heapify
    // parent( on node[i]) less than each child, recursively
    smallest := i
    l := 2 * i + 1
    r := 2 * i + 2
    //fmt.Printf("i: %d, l: %d, r: %d\n", i, l, r)

    if l < length && (*h.arr)[l] < (*h.arr)[smallest]{
        smallest = l
    }

    if r < length && (*h.arr)[r] < (*h.arr)[smallest]{
        smallest = r
    }

    if smallest != i {
        swap(h.arr,i, smallest)
        h.heapify(length, smallest)
    }

}


func heapSort(h *Heap){
    // max sort
    for i := h.size / 2 - 1; i >= 0 ; i-- {
        fmt.Printf("i: %d , size: %d\n", i, h.size)
        h.heapify(h.size, i)
        fmt.Println(h.arr)
    }

    for i := h.size - 1;i >= 0; i-- {
        swap(h.arr, 0, i)
        h.heapify(i, 0)
    }

}




func createHeap(cap int)(*Heap){
    heap_array := make([]int, cap)
    return&Heap{
        size : 0,
        capacity : cap,
        arr : &heap_array,
    }
}


func main(){
    heap := createHeap(11)

    heap.insertKey(3)
    heap.insertKey(2)
    heap.insertKey(1)
    heap.insertKey(15)
    heap.insertKey(5)
    heap.insertKey(4)
    heap.insertKey(45)
    fmt.Println(heap.arr)
    heapSort(heap)
    fmt.Println(heap.arr)

//    heap1 := createHeap(11)
//    heap1.insertKey(12)
//    heap1.insertKey(11)
//    heap1.insertKey(13)
//    heap1.insertKey(5)
//    heap1.insertKey(6)
//    heap1.insertKey(7)
//    fmt.Println(heap1.arr)
//    heapSort(heap1)
//    fmt.Println(heap1.arr)
}
