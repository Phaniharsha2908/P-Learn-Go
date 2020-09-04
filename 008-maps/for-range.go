package main

import "fmt"

func main() {
	eng := map[string]string{
		"up":   "above",
		"down": "below",
	}

	for k, v := range eng {
		fmt.Println("key=", k, "val=", v)
	}
}
