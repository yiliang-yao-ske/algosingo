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

    heap.heapify(i)

}


func swap(arr *[]int, x, y int){
    t := (*arr)[x]
    (*arr)[x] = (*arr)[y]
    (*arr)[y] = t
}


func (h *Heap)heapify(i int){

    smallest := i
    l := 2 * i + 1
    r := 2 * i + 2
    //fmt.Printf("i: %d, l: %d, r: %d\n", i, l, r)

    if l < h.size && (*h.arr)[l] < (*h.arr)[smallest]{
        smallest = l
    }

    if r < h.size && (*h.arr)[r] < (*h.arr)[smallest]{
        smallest = r
    }

    if smallest != i {
        swap(h.arr,i, smallest)
        h.heapify(smallest)
    }

}


func (h *Heap)sort(){
    // min sort
    for i := h.size / 2 - 1; i >= 0 ; i-- {
        h.heapify(i)
    }

    for i := h.size - 1;i >= 0; i-- {
        swap(h.arr, 0, i)
        h.heapify(0)
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
    fmt.Println(heap.arr)
    heap.insertKey(2)
    fmt.Println(heap.arr)
    heap.insertKey(1)
    fmt.Println(heap.arr)
    heap.insertKey(15)
    fmt.Println(heap.arr)
    heap.insertKey(5)
    fmt.Println(heap.arr)
    heap.insertKey(4)
    fmt.Println(heap.arr)
    heap.insertKey(45)
    fmt.Println(heap.arr)
    heap.sort()
    fmt.Println(heap.arr)
    // heap.minheapify(2)
    // fmt.Println(heap.arr)
    // fmt.Println(heap.extractMin())
    // fmt.Println(heap.arr)
    // fmt.Println(heap.getMin())
    // fmt.Println(heap.arr)
    // heap.decreaseKey(2, 1)
    // fmt.Println(heap.getMin())

}
