package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	balance float64
}
func (a *account)withdraw(amount float64, wg *sync.WaitGroup){
	defer wg.Done()
	if a.balance-amount<0{
		fmt.Printf("Balance not sufficient. Account balance is %v",a.balance)
		return
	}
	a.balance-=amount
	fmt.Printf("%v withrew from account.Account balance is %v\n ",amount, a.balance)
	time.Sleep(1)
}

func (a *account)deposit(amount float64, wg *sync.WaitGroup){
	defer wg.Done()
	a.balance=a.balance+amount
	fmt.Printf("%v deposited to account.Account balance is %v\n ",amount, a.balance)
	time.Sleep(1)
}

func main()  {
	var wg sync.WaitGroup
	a:=account{500}
	var mutex = &sync.Mutex{}
	wg.Add(1)
	mutex.Lock()
	go a.deposit(220.5,&wg)
	mutex.Unlock()

	wg.Add(1)
	mutex.Lock()
	go a.withdraw(400,&wg)
	mutex.Unlock()

	wg.Add(1)
	mutex.Lock()
	go a.withdraw(700.5,&wg)
	fmt.Println()
	mutex.Unlock()

	wg.Wait()

}
