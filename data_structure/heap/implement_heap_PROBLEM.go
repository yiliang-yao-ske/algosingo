package main


import (
    "fmt"
    "math"
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

func (heap *Heap)decreaseKey(i, new_value int) {
    (*heap.h_arr)[i] = new_value
    for i != 0 && (*heap.h_arr)[heap.parent(i)] > (*heap.h_arr)[i] {
        swap((*heap.h_arr)[i], (*heap.h_arr)[heap.parent(i)])
        i = heap.parent(i)
    }
 
}


func (heap *Heap)extractMin() int{
    if (heap.h_size <= 0){
        return math.MaxInt16
    }
    if heap.h_size == 1 {
        heap.h_size -= 1
        return (*heap.h_arr)[0]
    }

    root := (*heap.h_arr)[0]
    (*heap.h_arr)[0] = (*heap.h_arr)[heap.h_size -1]
    heap.h_size -= 1
    heap.minheapify(0)

    return root
}



func (heap *Heap)deleteKey(i int){
    heap.decreaseKey(i, math.MinInt16)
    heap.extractMin()
}


func (heap *Heap)minheapify(i int){
    l := heap.left(i)
    r := heap.right(i)
    // fmt.Printf("l index: %d  value: %d \n", l, (*heap.h_arr)[l])
    // fmt.Printf("r index: %d  value: %d \n", r, (*heap.h_arr)[r])
    smallest := i
    if l < heap.h_size && (*heap.h_arr)[l] < (*heap.h_arr)[i] {
        smallest = l
    }
    if r < heap.h_size && (*heap.h_arr)[r] < (*heap.h_arr)[smallest] {
        smallest = r
    }
    if smallest != i {
        swap((*heap.h_arr)[i], (*heap.h_arr)[smallest])
        heap.minheapify(smallest)
    }
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
