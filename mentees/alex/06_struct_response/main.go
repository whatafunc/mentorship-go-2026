package main

import (
	json "encoding/json"
	"fmt"
)

type Response struct {
	// поля с тегами
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	// код десериализации
	var res Response
	err := json.Unmarshal([]byte(rawResp), &res)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}

func main() {

	resp := `{
		"header": {
			"code": 0,
			"message": ""
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`
	res, err := ReadResponse(resp)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println(res)

	// Verify the data was unmarshaled correctly
	fmt.Printf("Email: %s\n", res.Data[0].Attributes.Email)
	fmt.Printf("Article IDs: %v\n", res.Data[0].Attributes.ArticleIds)
}
