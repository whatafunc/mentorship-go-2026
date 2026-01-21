package main

// Напишите функцию `Concat`, которая получает несколько слайсов
// и склеивает их в один длинный.
// { {1, 2, 3}, {4, 5}, {6, 7} }  => {1, 2, 3, 4, 5, 6, 7}

func Concat(slices [][]int) []int {
	res := []int{}
	for _, v := range slices {
		//fmt.Println(v)
		for _, sv := range v {
			res = append(res, sv)
		}

	}
	//fmt.Println(res)
	return res
}
