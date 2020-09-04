package main

import "fmt"

func main() {

	eng := map[string]string{
		"up": "above",
	}

	english := eng

	english["up"] = "down"

	fmt.Println(eng)
	fmt.Println(english)

}
