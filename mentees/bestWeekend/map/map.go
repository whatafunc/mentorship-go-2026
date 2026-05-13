package main

import (
	"fmt"
)

func outputAaa(v map[string]string) {
	println("outputAaa adres = ", v)
	println("outputAaa vocab slovo = ", v["aa"])

}

func main() {
	vocab := make(map[string]string, 1)
	fmt.Println(vocab)
	fmt.Println(vocab["aa"])
	if vocab["aa"] == "" {
		fmt.Println("key does not exist")
		vocab["aa"] = "valaaaa"
		fmt.Println("aa key in vocab = ", vocab["aa"])
		fmt.Println("FMT adres = ", &vocab)
		println("Print adres = ", vocab)
		outputAaa(vocab)
	}
	vocab["bb"] = "valbbbb"
	fmt.Println("Print adres = ", vocab)
	fmt.Println("FMT adres = ", &vocab)
	vocab["cb"] = "valbbbb"
	fmt.Println("FMT adres = ", &vocab)

}
