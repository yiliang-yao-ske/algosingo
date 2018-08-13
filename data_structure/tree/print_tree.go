package main


import (
    "fmt"
)


type Node struct{
    data int
    left *Node
    right *Node
}

func (node *Node)getDepth() int{
    if node == nil{return 0}

    lDepth := node.left.getDepth()
    rDepth := node.right.getDepth()
    if lDepth > rDepth{
        return lDepth + 1
    }else{
        return rDepth +1
    }
}


func newNode (data int) (node *Node){
    return &Node{
        data  : data,
        left  : nil,
        right : nil,
    }
}


func (node *Node)printPreorder(){
    if node == nil {return}

    node.left.printPreorder()
    node.right.printPreorder()

    fmt.Println(node.data)

}


func (node *Node)printInorder(){
    if node == nil {return}

    node.left.printInorder()
    fmt.Println(node.data)
    node.right.printInorder()

}


func (node *Node)printPostorder(){
    if node == nil {return}

    fmt.Println(node.data)
    node.left.printPostorder()
    node.right.printPostorder()

}


func (node *Node)printLevelorder(){

    h := node.height()

}

func main(){
    root := newNode(1)
    root.left = newNode(2)
    root.right = newNode(3)
    root.left.left = newNode(4)
    root.left.right = newNode(5)
    root.left.right.left = newNode(6)

    fmt.Println(root.data)
    fmt.Println("depth :",root.getDepth())

    root.printPreorder()
}
