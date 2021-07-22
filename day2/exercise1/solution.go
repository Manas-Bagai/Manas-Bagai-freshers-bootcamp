package main

import (
	"fmt"
	"sync"
)

type Result struct {
	result map[string]int
}
func (r Result) count(word string)  {
	for _,v:=range word{
		r.result[string(v)]+=1
	}
}

func main(){

	var mutex = &sync.Mutex{}
	a:=make(map[string]int)
	str:="abcdefghijklmnopqrstuvwxyz"
	for _,v:=range str{
		a[string(v)]=0
	}
	res:=Result{a}
	words:=[]string{"quick","brown","fox","lazy","dog"}
	for _,v2:=range words{
		mutex.Lock()
		go res.count(v2)
		mutex.Unlock()
	}
	fmt.Println(a)
}
