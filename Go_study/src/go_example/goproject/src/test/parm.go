package main

import (
	"fmt"
)

func test() {
	fmt.Println("test")
}

func list_elem(n int, resultChan chan []int) {
	var i int = 0
	var s = make([]int, 10, 20)
	for i = 0; i < n; i++ {
		fmt.Println(i)
		s[i] = i
	}
	resultChan <- s
}
func trytry() {
	resultChan := make(chan []int, 10)
	//list_elem(10, resultChan)
	go list_elem(10, resultChan)
	var my = <-resultChan
	my = append(my, 999)
	//for _, elem := range my {
	//	fmt.Println(elem)
	//}
}
