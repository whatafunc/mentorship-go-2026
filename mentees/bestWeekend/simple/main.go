package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func SetAge(p *Person, age int) {
	p.Age = age
}

func SetAgeNew(p *Person, age int) {
	p = &Person{Name: p.Name, Age: age} // Создаем новый экземпляр Person с обновленным возрастом
	fmt.Println("Inside SetAgeNew:", p)
}

func main() {
	Dima := &Person{Name: "Dima", Age: 20}
	println("Hello, World!", Dima.Name, Dima.Age)
	//Dima.Age = 21
	SetAge(Dima, 21)
	fmt.Println("Hello, World!", Dima)
	SetAgeNew(Dima, 22)
	fmt.Println("Hello, World!", Dima)

}
