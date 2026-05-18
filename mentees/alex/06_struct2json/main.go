package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Имайл адрес"`
	DateOfBirth time.Time
}

func main() {
	Coder1 := Person{
		Name:  "Alex",
		Email: "al@ex.com",
	}

	jsonData, err := json.Marshal(Coder1)
	if err != nil {
		fmt.Println("Error marshaling Coder1:", err)
		return
	}
	fmt.Println("jsonData for Coder1 // in []bytes = ", jsonData)

	fmt.Println("jsonData for Coder1 = ", string(jsonData))

	// ---

	req := struct {
		NameContains string `json:"name_contains"`
		Offset       int    `json:"offset"`
		Limit        int    `json:"limit"`
	}{
		NameContains: "Иван",
		Limit:        50,
	}

	reqRaw, _ := json.Marshal(req)
	fmt.Println(string(reqRaw))

}
