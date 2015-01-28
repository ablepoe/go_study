// copyTest
package main

import (
	"fmt"
)

func copyTest() {
	s := [5]int{1, 2, 3, 4, 5}

	s1 := s[1:3]

	//s2 := make([]int, 2) // make已经给s2分配好底层array的空间，并且用0填补array
	//fmt.Println(s2)

	s2 := s[2:4]
	fmt.Println("before copy s1 is:", s1, "s2 before copy is:", s2)
	fmt.Println("before copy s is:", s)
	copy(s2, s1)
	fmt.Println("after copy s1 is:", s1, "after copy s2 is", s2)
	fmt.Println("after copy s is:", s)
	s2[0] = 9
	s1[0] = 99

	fmt.Println(s1, s2)

}
