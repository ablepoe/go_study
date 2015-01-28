package main

import (
	"fmt"
)

func main() {
	p := Student{Person{2, "张三"}, 25}
	p.test()

}

type Person struct {
	Id   int
	Name string
}

type Student struct {
	Person
	Score int
}

func (this Person) test() {
	fmt.Println("person test")
}

func (this Student) test() {
	fmt.Println("student test")
	this.Person.Id = 5
	fmt.Println("Person id is:", this.Id)
	this.Id = 10
	fmt.Println("Person id is:", this.Id)
}
