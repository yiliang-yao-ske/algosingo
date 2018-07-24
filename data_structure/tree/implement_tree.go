package main

import (
	"fmt"
)
/* 
	Interface for binary tree in go
-------------------------------------

	1) getHeight
	2) getSize
	3) #printPreorder
	4) #printPostorder
	5) printInorder




*/

type Node struct{
	data int
	left *Node
	right *Node
}

func newNode(data int) (node *Node){

	return &Node{
			data : data,
			left : nil,
			right: nil,
		}

}



func (node *Node)getHeight() int{
    if node == nil{return 0}

    lDepth := node.left.getHeight()
    rDepth := node.right.getHeight()
    if lDepth > rDepth{
        return lDepth + 1
    }else{
        return rDepth +1
    }
}


func (node *Node)getSize() int{
	// return nodes number include root
	if node == nil { return 0 }

	lSize := node.left.getSize()
	rSize := node.right.getSize()

	return lSize + rSize + 1

}


func main() {
	
	root := newNode(1)

	fmt.Println("root value: ", root.data)

	root.left = newNode(2)
	root.right = newNode(3)

	root.left.left = newNode(4)

	root.right.left = newNode(6)

	fmt.Println("root.left value: ", root.left.data)
	fmt.Println("root.left.left value: ", root.left.left.data)


	fmt.Println("root Height: ", root.getHeight())

	fmt.Println("root Size: ", root.getSize())	

	return
}