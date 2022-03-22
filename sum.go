package main

import "fmt"

func sum(numbers ...int) (plus int) {
	for _, val := range numbers {
		plus = plus + val
	}
	return
}

func main() {
	fmt.Println(sum(2, 4, 6))
}
