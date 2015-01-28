package main

import (
	"fmt"
	my "myImport"
)

func set_value(x_ptr *int) {
	*x_ptr = 100
}
func main() {

	x_ptr := new(int)
	fmt.Println(x_ptr)
	set_value(x_ptr)
	//x_ptr指向的地址
	fmt.Println(x_ptr)
	//x_ptr本身的地址
	fmt.Println(&x_ptr)
	//x_ptr指向的地址值
	fmt.Println(*x_ptr)

	test()
	trytry()
	my.TestImport()
}
