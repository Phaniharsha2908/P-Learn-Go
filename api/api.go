package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting our app")

	response, err := http.Get("https://api.coinbase.com/v2/prices/BTC-USD/buy")
	if err != nil {
		log.Fatal("The Http request failed with an error", err)
	}

	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))

}
