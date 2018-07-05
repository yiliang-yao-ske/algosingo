package main

import "fmt"


type Node struct{
    data int
    next *Node
}

type LinkedList struct{
    head *Node
}

func printList(list *LinkedList){
    node := list.head
    for node != nil {
        fmt.Printf(" %d ", node.data)
        node = node.next
    }
    fmt.Println("")
}

func (list *LinkedList)push(value int){

    new_node := &Node{}

    new_node.data = value
    new_node.next = list.head

    list.head = new_node

}


func (list *LinkedList)append(value int){


    new_node := &Node{
        data : value,
        next : nil,
    }

    if list.head == nil {
        list.head = new_node
        return
    }

    last := list.head
    for last.next != nil {
        last = last.next
    }
    last.next = new_node
}

func (list *LinkedList)insertAfter(prev_node *Node, value int){
    if prev_node == nil {
        fmt.Printf("The given previous node must in LinkedList")
        return
    }

    new_node := &Node{
        data : value,
        next : prev_node.next,
    }
    prev_node.next = new_node
}

func main(){

    list := &LinkedList{}

    list.append(2)

    list.append(3)

    list.push(-1)

    list.insertAfter(list.head.next, 0)

    printList(list)

    return
}
