package main

func main() {
	/*
		Дан массив целых чисел A и целое число k.
		Нужно найти и вывести индексы пары чисел, сумма которых равна k.
		Если таких чисел нет, то вернуть пустой слайс.
		Индексы можно вернуть в любом порядке.
	*/

	Arr := [10]int{33, 65, 81, 233, 3, 6, 5, 8, 1, 23}
	k := 34
	resMap := make(map[int]int)
	for i, val := range Arr { // 33, 65, 81, 233, 3, 6, 5, 8, 1, 23
		println("=start=", i, val)

		diff := k - val
		if _, ok := resMap[val]; !ok {
			resMap[val] = i
			println("Added to map:", val, i)
		}

		if idx, ok := resMap[diff]; ok {
			println("Indexes:", idx, i)
			break
		}

	}
}
