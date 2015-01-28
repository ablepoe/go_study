package main

//import "fmt"

//func main() {
//	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	s2 := make([]int, 3, 20)
//	var n int
//	n = copy(s2, s1)
//	//fmt.Println(n, s2, len(s2), cap(s2))
//	fmt.Println("s1 first time is ", s1, "1")

//	s3 := s1[4:6]
//	//s1{4,5}
//	//s3{4,5}
//	fmt.Println("=======", n, s3, len(s3), cap(s3))
//	fmt.Println("s3 is ", s3, "s1 second time is", s1, "2")

//	s3 = []int{7, 8}
//	//n = copy(s3, s1[1:5])
//	//s1{1,2,3,4}
//	//s3{1,2}
//	//fmt.Println(n, s3, len(s3), cap(s3))

//	fmt.Println("s3 is ", s3, "s1 third time is", s1, "3")

//	fmt.Println("s1[1:6] is", s1[1:6], "s1[1:6] len is ", len(s1[1:6]), "s1[1:6] cap is", cap(s1[1:6]))
//	fmt.Println("s1[4:5] is", s1[4:5], "s1[4:5] len is ", len(s1[4:5]), "s1[4:5] cap is", cap(s1[4:5]))
//	fmt.Println("s1[4:6] is", s1[4:6], "s1[4:6] len is ", len(s1[4:6]), "s1[4:6] cap is", cap(s1[4:6]))
//	fmt.Println("s1[4:7] is", s1[4:7], "s1[4:7] len is ", len(s1[4:7]), "s1[4:7] cap is", cap(s1[4:7]))
//	fmt.Println("s1[4:8] is", s1[4:8], "s1[4:8] len is ", len(s1[4:8]), "s1[4:8] cap is", cap(s1[4:8]))
//	//fmt.Println(cap(s1))

//}
func main() {
	copyTest()
}
