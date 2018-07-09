package main

import (
    "fmt"
    "math"
)


type Stack struct {
    top int
    capacity uint16
    array *[]int
}




func createStack(capa uint16) (stack *Stack){

    arr := make([]int, capa)
    return &Stack{
        top : -1,
        capacity : capa,
        array : &arr,
    }

}

func (stack *Stack)isEmpty() bool{
    if stack.top == -1{
        return true
    }else{
        return false
    }
}

func (stack *Stack)isFull() bool{
    if uint16(stack.top) == stack.capacity {
        return true
    }else{
        return false
    }
}


func (stack *Stack)push(item int){
    if stack.isFull(){
        return
    }
    stack.top += 1
    (*stack.array)[stack.top] = item
    fmt.Printf("%d pushed to stack \n", item)
}


func (stack *Stack)pop() int{
    if stack.isEmpty(){
        return math.MinInt16
    }
    item := (*stack.array)[stack.top]
    stack.top -= 1
    return item
}


type Queue struct {
    stack1 *Stack
    stack2 *Stack
}

func createQueue(capacity uint16) (queue *Queue){
    stack1 := createStack(capacity)
    stack2 := createStack(capacity)


    return &Queue{
        stack1 : stack1,
        stack2 : stack2,
    }
}


func (queue *Queue)enQueue(item int){
    if queue.stack1.isFull() == false{
        queue.stack1.push(item)
    }else{
        fmt.Println("Enqueue is full")
    }

}

func (queue *Queue)deQueue() (item int){
    if queue.stack1.isEmpty() && queue.stack2.isEmpty(){
        fmt.Println("Error for dequeue")
        return math.MinInt16
    }
    if queue.stack2.isEmpty() {
        for queue.stack1.isEmpty() == false {
            queue.stack2.push(queue.stack1.pop())
        }
    }

    return queue.stack2.pop()
}



func main() {

    queue1 := createQueue(10)

    fmt.Printf("Stack is empty: %t \n", queue1.stack1.isEmpty())

    fmt.Println("THIS SHALL BE error: ",queue1.deQueue())

    queue1.enQueue(1)

    queue1.enQueue(2)

    queue1.enQueue(3)

    fmt.Printf("dequeue for : %d\n", queue1.deQueue())
    fmt.Printf("dequeue for : %d\n", queue1.deQueue())

//    stack1.push(1)
//
//    stack1.push(2)
//
//    fmt.Printf("%d poped out\n", stack1.pop())
//    fmt.Printf("%d poped out\n", stack1.pop())
//    fmt.Printf("%d poped out\n", stack1.pop())

}
