package main

import "fmt"

//人的结构体
type Person struct {
	name string
	age  int
}
type Employee struct {
	EmployeeID int
	Person
}

func (e Employee) PrintInfo() {
	fmt.Printf("EmployeeID: %d, Name: %s, Age: %d\n", e.EmployeeID, e.name, e.age)
}
func main() {
	p := Person{"zhangsan", 20}
	e := Employee{001, p}
	e.PrintInfo()

}
