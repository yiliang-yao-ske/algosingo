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

func (list *LinkedList)reverse(){
    var prev *Node = nil
    curr := list.head

    for curr != nil {
        next := curr.next
        curr.next = prev
        prev = curr
        curr = next
    list.head = prev
    }

}


func main(){

    list := &LinkedList{}

    list.append(2)

    list.append(3)

    list.push(-1)

    printList(list)
    list.insertAfter(list.head.next, 0)

    printList(list)

    list.reverse()
    printList(list)

    return
}
