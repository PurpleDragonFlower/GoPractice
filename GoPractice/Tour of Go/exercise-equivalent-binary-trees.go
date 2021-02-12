package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	
	ch <- t.Value
	
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	Walk(t1, ch1)
	Walk(t2, ch2)
	for i:= 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Testing Walk Functionality")
	ch := make(chan int,10)
	go Walk(tree.New(1), ch)
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Walk concluded")
	
	fmt.Println("Testing Same Tree Functionality")
	t1,t2 := tree.New(1), tree.New(2)
	
	fmt.Println("t1: %v", t1.String())
	fmt.Println("t2: %v", t2.String())
	
	fmt.Println("Same?: t1==t1 %v", Same(t1,t1))
	fmt.Println("Same?: t1==t2 %v", Same(t1,t2))
	fmt.Println("Same Tree concluded")
}
