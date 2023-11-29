package restclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SampleRestClientBuiltIn() {
	URI := "https://jsonplaceholder.typicode.com/todos/1"
	req, err := http.NewRequest(http.MethodGet, URI, nil)
	if err != nil {
		log.Println("err:", err.Error())
	}

	// Simulasi request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err:", err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("err:", err.Error())
	}

	fmt.Println("resp.body:", string(body))
}
