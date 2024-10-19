package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		num := i % 3
		if num == 0 {
			fmt.Println(i)
		}
	}

}
