package main

import (
	"fmt"
	"math/rand"
	"sync"
)


func getGrade(wg *sync.WaitGroup, grade chan<-int){
	defer wg.Done()
	a:=rand.Intn(10)
	grade<-a
}

func main(){
	var wg sync.WaitGroup
	grade:=make(chan int)
	numStudents:=200
	final:=make(map[int]int)
	for i:=1;i<=numStudents;i++{
		c:=rand.Intn(5)
		wg.Add(c)
		go getGrade(&wg,grade)
		v:=<-grade
		final[i]=v
	}
	fmt.Println(final)
}
