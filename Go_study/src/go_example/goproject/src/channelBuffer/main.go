package main

import (
	"fmt"
)

var a string
var c = make(chan int, 1) //此处定义为缓冲的channel
func ff() {
	a = "hello world"
	ca := <-c
	fmt.Println(ca, "99")
}

func main() {
	go ff() //开启一个goroutine
	c <- 3
	fmt.Println(a, "3")
	b := <-c
	c <- 2
	fmt.Println(a, "2")
	b = <-c
	c <- 1
	fmt.Println(a, "1")
	b = <-c
	c <- 0
	fmt.Println(a, "0")

	fmt.Print(b)
	//var input string
	//fmt.Scanln(&input)
}
