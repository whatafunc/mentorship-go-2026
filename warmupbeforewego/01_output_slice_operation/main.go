package main

import (
	"fmt"
)

var a = 1 // <- уровень пакета

func main() {
	fmt.Println("1: ", a) ///1

	a := 2                // <-- уровень блока функции
	fmt.Println("2: ", a) ///2
	{
		a := 3                // <-- уровень пустого блока
		fmt.Println("3: ", a) ///3
	}
	fmt.Println("4: ", a) ///2

	f()
}

func f() {
	fmt.Println("5: ", a) ///1
}
