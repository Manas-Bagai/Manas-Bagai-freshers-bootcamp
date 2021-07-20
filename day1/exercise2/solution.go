package main

type Tree struct{
	Left *Tree
	Right *Tree
	Val int
}

func NodetoTree(tree *Tree, node *Tree){
	if tree.Val== 1{
		tree.Val=node.Val
		return
	}

}

func TreeMak(values []int)( *Tree){
	root:= Tree{Left: nil, Right: nil, Val: 1}
	for _,v := range values{
		node:= Tree{Val: v}
		NodetoTree(&root,&node)
	}
	return &root
}
