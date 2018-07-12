package main

import (
    "fmt"
    "math"
)


type Stack struct {
    top int
    capacity uint16
    array *[]int
    maxEle int
}


func createStack(capa uint16) (stack *Stack){

    arr := make([]int, capa)
    return &Stack{
        top : -1,
        capacity : capa,
        array : &arr,
        maxEle : math.MaxInt16,
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

//func (stack *Stack)peek() (item int){
//    if stack.isEmpty() {
//        fmt.Println("Stack is empty")
//        return math.MinInt16
//    }
//    
//}


func (stack *Stack)push(item int){
    if stack.isFull(){
        return
    }

    if stack.isEmpty(){
        stack.top += 1
        (*stack.array)[stack.top] = item
        stack.maxEle = item
        fmt.Printf("%d pushed to stack \n", item)
        fmt.Printf("%d update to maxEle \n", stack.maxEle)
    }else if (item <= stack.maxEle){
        stack.top += 1
        (*stack.array)[stack.top] = item
        fmt.Printf("%d pushed to stack \n", item)
        fmt.Printf("%d update to maxEle \n", stack.maxEle)
    }else{
        stack.top += 1
        (*stack.array)[stack.top] = (item + stack.maxEle)
        stack.maxEle = item
        fmt.Printf("%d pushed to stack \n", (*stack.array)[stack.top])
        fmt.Printf("%d update to maxEle \n", stack.maxEle)
    }


}


func (stack *Stack)pop() int{
    if stack.isEmpty(){
        return math.MinInt16
    }
    item := (*stack.array)[stack.top]
    fmt.Printf("item: %d\n", item)
    stack.top -= 1

    if item > stack.maxEle {
        fmt.Printf("min: %d\n", stack.maxEle)
        stack.maxEle = (item - stack.maxEle)
        fmt.Printf("min: %d\n", stack.maxEle)
        return item - stack.maxEle
    }

    return item
}


func main() {

    stack1 := createStack(10)

    fmt.Printf("%d %p\n", stack1.top, stack1.array)
    fmt.Printf("Stack is empty: %t \n", stack1.isEmpty())


    stack1.push(3)
    stack1.push(5)
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    stack1.push(6)
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    stack1.push(1)
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    stack1.push(19)
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    stack1.push(-1)
    fmt.Printf("%d is maxEle\n", stack1.maxEle)

    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d poped out\n", stack1.pop())
    fmt.Printf("%d is maxEle\n", stack1.maxEle)
    fmt.Printf("%d poped out\n", stack1.pop())

}
