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


func main() {

    stack1 := createStack(10)

    fmt.Printf("%d %p\n", stack1.top, stack1.array)
    fmt.Printf("Stack is empty: %t \n", stack1.isEmpty())


    stack1.push(1)

    stack1.push(2)

    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d poped out\n", stack1.pop())

}
