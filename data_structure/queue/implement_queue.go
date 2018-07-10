package main

import (
    "fmt"
    "math"
)

type Queue struct{
    front int
    rear int
    size int
    capacity int
    array *[]int

}


func createQueue(capacity int) (queue *Queue){
    arr := make([]int, capacity)
    return &Queue{
        front : 0,
        size : 0,
        rear : capacity -1,
        capacity : capacity,
        array : &arr,
    }
}

func (queue *Queue)isFull() bool{

    return (queue.size == queue.capacity)
}

func (queue *Queue)isEmpty() bool{
    return (queue.size == 0)
}

func (queue *Queue)enqueue(item int){
    if queue.isFull(){
        return
    }
    queue.rear = (queue.rear + 1)% queue.capacity
    (*queue.array)[queue.rear] = item
    queue.size += 1
    fmt.Println(item, " enqueued to queue")
}

func (queue *Queue)dequeue()(item int){
    if queue.isEmpty(){
        fmt.Println("queue is empty")
        return math.MinInt16
    }
    item = (*queue.array)[queue.front] % queue.capacity
    queue.front += 1
    queue.size -= 1
    return item

}


func main(){

    queue := createQueue(30)

    fmt.Printf("Create queue at: %p\n", queue)

    fmt.Printf("At the beginning, the queue is empty %t\n", queue.isEmpty())


    queue.enqueue(1)
    queue.enqueue(2)
    queue.enqueue(3)
    fmt.Printf("%d\n",queue.dequeue())
    fmt.Printf("%d\n",queue.dequeue())
    fmt.Printf("%d\n",queue.dequeue())
    fmt.Printf("%d\n",queue.dequeue())
    fmt.Printf("%d\n",queue.dequeue())

}
