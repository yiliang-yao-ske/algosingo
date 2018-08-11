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

/* 
	below are for BST
*/

func (node *Node)insert(key int) (n *Node){
	if node == nil { return newNode(key)}

	if (key < node.data){
		node.left = node.left.insert(key)
	}else if (key > node.data){
		node.right = node.right.insert(key)
	}
	return node
}


func printInorder(root *Node){
	if root != nil{
		printInorder(root.left)
		fmt.Println(root.data)
		printInorder(root.right)
	}
}

func search(root *Node, key int)(node *Node){
	if root == nil || root.data == key {
		return root
	}

	if root.data < key {
		return search(root.right, key)
	}
	return search(root.left, key)
}


func main() {
	
	root := newNode(50)

	fmt.Println("root value: ", root.data)

	// root.left = newNode(2)
	// root.right = newNode(3)

	// root.left.left = newNode(4)

	// root.right.left = newNode(6)

	root.insert(30)
	root.insert(20)
	root.insert(40)
	root.insert(70)
	root.insert(60)
	root.insert(80)
	//To redundant value, we may check for this solution
	//https://www.geeksforgeeks.org/how-to-handle-duplicates-in-binary-search-tree/
	// add counter for each value
	//root.insert(20)	


	fmt.Println("root.left value: ", root.left.data)
	fmt.Println("root.left.left value: ", root.left.left.data)
	fmt.Println("root.left.right value: ", root.left.right.data)


	fmt.Println("root Height: ", root.getHeight())

	fmt.Println("root Size: ", root.getSize())	
	printInorder(root)

	fmt.Println("find value with: ", search(root, 50))	


	return
}