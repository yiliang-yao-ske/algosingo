package main

import "fmt"


type Node struct{
    data int
    next *Node
}

func printList(node *Node){
    for node != nil {
        fmt.Printf(" %d ", node.data)
        node = node.next
    }
    fmt.Println("")
}


func main(){

    var head Node
    var second Node
    var third Node

    head.data = 1
    head.next = &second

    second.data = 2
    second.next = &third

    third.data = 3
    third.next = nil

    printList(&head)

    return
}
