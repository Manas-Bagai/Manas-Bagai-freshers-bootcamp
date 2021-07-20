package main

import "fmt"

type Tree struct{
	Left *Tree
	Right *Tree
	Val string
}

func (t *Tree) preorder() {
	if t != nil {
		fmt.Printf("%v", t.Val)
		t.Left.preorder()
		t.Right.preorder()
	}
}
func (t *Tree) postorder(){
	if t != nil {
		t.Left.preorder()
		t.Right.preorder()
		fmt.Printf("%v", t.Val)
	}
}

func main(){
	node1:=Tree{Left:nil, Right: nil, Val: "a"}
	node2:= Tree{Left:nil, Right: nil, Val: "b"}
	node3:= Tree{Left:nil, Right: nil, Val: "c"}
	node4:= Tree{Left:&node2, Right: &node3, Val: "-"}
	node5:= Tree{Left:&node1, Right: &node4, Val: "+"}
	node5.preorder()
	fmt.Println(" ")
	node5.postorder()
}
