package main

import (
    "fmt"
    "reflect"
)


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

func (list *LinkedList)deleteNode(key int){
    if list.head == nil {
        fmt.Printf("List is empty\n")
        return
    }

    temp := list.head
    if temp.data == key{
        list.head = temp.next
        //gc temp node
        return
    }


    for temp.next != nil && temp.next.data != key {
        temp = temp.next
    }

    if temp.next == nil {return}

    temp.next = temp.next.next

    //gc temp.next

}



func (list *LinkedList)insertAfter(prev_node *Node, value int){
    if prev_node == nil {
        fmt.Printf("The given previous node must in LinkedList\n")
        return
    }

    new_node := &Node{
        data : value,
        next : prev_node.next,
    }
    prev_node.next = new_node
}

func (list *LinkedList)isLoop() bool {
    slowNode := list.head
    fastNode := list.head


    for slowNode != nil && fastNode != nil && fastNode.next != nil{
        fmt.Printf("fastNode: %p\n", fastNode)
        fmt.Printf("slowNode: %p\n", slowNode)
        slowNode = slowNode.next
        fastNode = fastNode.next.next
        if fastNode == slowNode {
            return true
        }
    }

    return false
}



func main(){

    list := &LinkedList{}

    list.append(1)

    list.append(2)

    list.append(3)

    list.append(4)

    list.append(5)

    printList(list)

    list.head.next.next.next.next.next = list.head.next.next

    fmt.Printf("%t \n",list.isLoop())


    fmt.Println(reflect.TypeOf(true))

    return
}
