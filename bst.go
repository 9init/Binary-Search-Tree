package main
import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct{
	left *Node
	right *Node
	value int
	child bool
}

func insert(root *Node, v int){
	if root==(&Node{}){
		*root=Node{nil,nil,v,true}
	}else if v > root.value { //right node
		if root.right != nil {insert(root.right, v)
		}else{root.right=&Node{nil,nil,v,true}}
	}else if v < root.value { //left node
		if root.left != nil {
			insert(root.left, v)
		}else{root.left=&Node{nil,nil,v,true}}
    	}
}

func traverse(root *Node) {
	if root!=nil{
		if root.left != nil {traverse(root.left)}
//		fmt.Print(root.value," ")
		if root.right != nil {traverse(root.right)}
	}
}

func main() {
	root:=new(Node)
	const SIZE = 900000
	var a [SIZE]int

	fmt.Printf("Generating random array with %v values...\n", SIZE)

	start := time.Now()

	for i := 0; i < SIZE; i++ {
		a[i] = rand.Intn(SIZE)
	}

	end := time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
	fmt.Printf("Filling the tree with %v nodes...\n", SIZE)

	start = time.Now()

	for i := 0; i < SIZE; i++ {
		insert(root, a[i])
	}

	end = time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
	fmt.Printf("Traversing all %v nodes in tree...\n", SIZE)

	start = time.Now()

	traverse(root)

	end = time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
}
