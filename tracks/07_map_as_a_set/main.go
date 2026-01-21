package main

import "fmt"

func main() {

	// Создание множества
	set := make(map[int]struct{})

	// Добавление элементов
	set[103] = struct{}{}
	set[42] = struct{}{}
	set[777] = struct{}{}

	// Проверка наличия элемента
	if _, exists := set[103]; exists {
		fmt.Println("103 есть в множестве")
	}

	// Удаление элемента
	delete(set, 42)
	fmt.Println("После удаления 42:", set)

	// Итерация по множеству
	fmt.Println("Элементы множества:")
	for key := range set {
		fmt.Println(key)
	}
}
