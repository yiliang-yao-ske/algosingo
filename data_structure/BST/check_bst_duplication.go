package main

import (
	"fmt"
	"sync"
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


func (node *Node)minValueNode()*Node{
	current := node.right

	for current.left != nil{
		current = current.left
	}
	return current
}

func deleteNode(root *Node, key int)*Node{

	if root == nil {return root}

	if key < root.data{
		root.left = deleteNode(root.left, key)
	}else if key > root.data {
		root.right = deleteNode(root.right, key)
	}else{	//find correct node to delete, then

		// if node has no child
		if root.left == nil && root.right == nil{
			return nil
		}else if root.left == nil{
			return root.right
		}else if root.right == nil{
			return root.left
		}
		// find smallest node in bigger value side(right)
		temp := root.minValueNode()

		root.data = temp.data

		root.right = deleteNode(root.right, temp.data)
	}
	return root
}


/* Implement set*/

type Set struct {
	m map[int]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m: map[int]bool{},
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := []int{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}






/* Check if duplicate */
func checkDupUtil(root *Node, set *Set)bool{
	if root == nil {return false}

	if set.Has(root.data){
		return true
	}

	set.Add(root.data)

	return checkDupUtil(root.left, set) ||
		   checkDupUtil(root.right, set)

}



func (root *Node)checkDup()(bool){

	set := NewSet()

	return checkDupUtil(root, set)

}



func main() {
	
	root := newNode(1)

	fmt.Println("root value: ", root.data)

	root.left = newNode(2)
	root.right = newNode(3)

	root.left.left = newNode(2)





	fmt.Println("root.left value: ", root.left.data)
	fmt.Println("root.right value: ", root.right.data)
	fmt.Println("root.left.left value: ", root.left.left.data)



	fmt.Println("root Height: ", root.getHeight())

	fmt.Println("root Size: ", root.getSize())	
	printInorder(root)


	fmt.Println(root.checkDup())



	return
}