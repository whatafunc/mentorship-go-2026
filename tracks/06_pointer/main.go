package main

import "fmt"

func main() {
	//
	var a int = 10
	var p *int = &a // p is a pointer to an integer, storing the address of a

	fmt.Println("Value of a:", a)                // Output: Value of a: 10
	fmt.Println("Address of a:", &a)             // Output: Address of a: <address>
	fmt.Println("Value of p (address of a):", p) // Output: Value of p (address of a): <address>
	fmt.Println("Value pointed to by p:", *p)    // Output: Value pointed to by p: 10

	*p = 20 // Changing the value at the address stored in p

	fmt.Println("New value of a after changing *p:", a) // Output: New value of a after changing *p: 20
	// Demonstrating pointer as function argument
	b := 30
	fmt.Println("b", b)
	modifyValue(&b)
	fmt.Println("b after modifyValue:", b)
}

func modifyValue(res *int) {
	fmt.Println("result = ", res)      // should be address
	fmt.Println("value of b = ", *res) // should be 30
	*res = 40
}
