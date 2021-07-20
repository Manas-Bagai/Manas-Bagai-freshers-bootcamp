package main

import "fmt"

type salary interface{
	FindSalary() int
}

type FullTime struct {
	basic int
}

type Contractor struct {
	basic int
}

type Freelancer struct{
	HoursPerDay int
	basicHourly int
}

func (f FullTime) FindSalary() int{
	return f.basic*28
}

func (c Contractor) FindSalary() int{
	return c.basic*28
}

func (fr Freelancer) FindSalary() int{
	return fr.basicHourly*fr.HoursPerDay*28
}

func main(){
	emp1:= FullTime{basic: 500}
	fmt.Println(emp1.FindSalary())
	emp2:=Contractor{basic:100}
	fmt.Println(emp2.FindSalary())
	emp3:=Freelancer{HoursPerDay: 20, basicHourly: 10}
	var a salary
	a=emp3
	fmt.Println(a.FindSalary())
}
