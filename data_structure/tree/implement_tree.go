package main


import (
    "fmt"
)

type Tree interface{
//    getElem()
//    setElem()
//    getParent()
//    getFirstChild()
//    getNextSibling()
//    getSize()
//    getHeight()
    getDepth()
}


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



func maxDepth(node *Node) int{

    if node == nil {
        return 0
    }

    lDepth := maxDepth(node.left)
    rDepth := maxDepth(node.right)
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


func main(){
    root := newNode(1)
    root.left = newNode(2)
    root.right = newNode(3)
    root.left.left = newNode(4)
    root.left.right = newNode(5)
    root.left.right.left = newNode(6)

    fmt.Println(root.data)
    //fmt.Println(maxDepth(root))
    fmt.Println("depth :",root.getDepth())

}
