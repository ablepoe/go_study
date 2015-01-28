package main

import (
	"fmt"
	"time"
)

var i int = 0
var j int = 0

func fixed_shooting(msg_chan chan string) {
	for {
		msg_chan <- "fixed shooting"
		fmt.Println("continue fixed shooting...")
		i++
	}
}

func count(msg_chan chan string) {
	for {
		msg := <-msg_chan
		fmt.Println(msg, "i is ", i, "j is ", j)
		time.Sleep(time.Second * 1)
	}
}

func three_point_shooting(msg_chan chan string) {
	for {
		msg_chan <- "three point shooting"
		fmt.Println("continue three point shooting...")
		j++
	}
}

func doubleCount(msg_chan1 chan string, msg_chan2 chan string) {
	for {
		select {
		case msg1 := <-msg_chan1:
			fmt.Println(msg1, "i is ", i, "j is ", j)
		case msg2 := <-msg_chan2:
			fmt.Println(msg2, "i is ", i, "j is ", j)
		}
	}
}

func main() {
	var c chan string
	c = make(chan string)
	var c2 chan string
	c2 = make(chan string)

	go fixed_shooting(c)
	go three_point_shooting(c2)
	//go count(c)
	//go doubleCount(c, c2)
	go func() {
		for {
			select {
			case msg1 := <-c:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}

	}()

	var input string
	fmt.Scanln(&input)
}
