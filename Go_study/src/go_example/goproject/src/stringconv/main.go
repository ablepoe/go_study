// stringconv project main.go
package main

import (
	"flag"
	"fmt"
	"strconv"
)

var i *int = flag.Int("p", 1111, "Port to listen.")

func lalala() {
	str := strconv.Itoa(*i)
	fmt.Println(str)
}

func main() {
	//str := strconv.Itoa(*i)
	//fmt.Println(str)
	lalala()
}
