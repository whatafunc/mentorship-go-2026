package main

import "fmt"

func main() {

	sl := []int8{1, 2, 3, 4, 5}
	println("simmple print slice addres", &sl)

	sl = append(sl, 6)
	//sl[101] = 10 //panic: runtime error: index out of range [101] with length 6

	fmt.Println("fmt print slice", sl)
	println("simmple print slice is", sl)
	println("simmple print slice addres", &sl)
	fmt.Println("addrs of slice", &sl)
	addr := &sl
	println("simmple print slice addr = ", addr)
	println("simmple print of slice = ", *addr)
	fmt.Println("slice is", *addr)

	// ----------------------------------------------
	fmt.Println("-----------------")
	sl = append(sl, 60, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	println("simmple print slice is", sl)
	println("simmple print slice addres", &sl)

	// ----------------------------------------------
	fmt.Println("-----------------")
	arr2 := [5]int{}
	arr2[0] = 1
	fmt.Println(arr2)
	if arr2[3] == 0 {
		fmt.Println("arr2[3] is zero")
		arr2[3] = 400
		fmt.Println("arr2[3] is now", arr2[3])
		//arr2[5] = 0 //out of bound.
	}
}
