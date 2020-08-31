package main

import "fmt"

func main() {

	b := [4]int{43, 65, 87, 89}

	for i, v := range b {
		fmt.Println(i, v)

	}
	for _, v := range b { // ignore index
		fmt.Println(v)

	}
	for i := range b { //ignore value but _ not necessary
		fmt.Println(i)

	}
	for i := 0; i < len(b); i++ {
		fmt.Println(i, b[i])
	}

}
