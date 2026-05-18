package main

import (
	"fmt"
)

func outputAaa(v map[string]string) {
	println("outputAaa adres = ", v)
	println("outputAaa vocab slovo = ", v["aa"])

}

// Напишите функцию, которая убирает дубликаты, сохраняя порядок слайса:
func RemoveDuplicates(input []string) []string {
	seen := make(map[string]bool, 1)
	result := make([]string, 0)

	for _, item := range input {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result

}

func main() {
	vocab := make(map[string]string, 1)
	fmt.Println(vocab)
	fmt.Println(vocab["aa"])
	if vocab["aa"] == "" {
		fmt.Println("key does not exist")
		vocab["aa"] = "valaaaa"
		fmt.Println("aa key in vocab = ", vocab["aa"])
		fmt.Println("Our map = ", &vocab)
		println("Print adres = ", vocab)
		outputAaa(vocab)
	}
	vocab["bb"] = "valbbbb"
	fmt.Println("Print adres = ", vocab)
	fmt.Println("Map = ", &vocab)
	vocab["cb"] = "valbbbb"
	fmt.Println("Map = ", &vocab)

	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	Cleanresults := RemoveDuplicates(input)
	fmt.Println("Cleanresults // no duplicates = ", Cleanresults)
	fmt.Println("Number of NO duplicates = ", len(Cleanresults))
	for i, res := range Cleanresults {
		fmt.Printf("Cleanresults[%d] = %s\n", i, res)
	}
}
