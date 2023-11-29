package restclient

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func SampleRestClientResty() {
	client := resty.New()
	URI := "https://jsonplaceholder.typicode.com/todos"
	resp, err := client.R().Get(URI)

	if err != nil {
		log.Println("err:", err.Error())
	}

	var todos []Todo
	err = json.Unmarshal(resp.Body(), &todos)
	if err != nil {
		log.Println("err:", err.Error())
	}

	// rumus pagination
	page := 2
	size := 20

	startIndex := (page - 1) * size
	endIndex := size + startIndex

	// cek kondisi jika page melebihi length
	if endIndex > len(todos) {
		endIndex = len(todos)
	}

	fmt.Println("page:", page)
	fmt.Println("size:", size)
	fmt.Println("startIndex:", startIndex)
	fmt.Println("endIndex:", endIndex)
	fmt.Println(todos[startIndex:endIndex])
}
