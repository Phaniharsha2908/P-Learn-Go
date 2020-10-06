package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting our app")
	jsonData := map[string]string{"firstname": "Jon", "lastname": "Peter"}
	jsonValue, _ := json.Marshal(jsonData)

	//response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))

	request, _ := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatal("The Http request failed with an error", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

}
