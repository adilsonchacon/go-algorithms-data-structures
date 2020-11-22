package main

import (
	"fmt"
)

func sort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		j := i - 1
		for j >= 0 {
			if numbers[j] > numbers[j+1] {
				aux := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = aux
				j--
			} else {
				break
			}
		}
	}
}

func main() {
	var numbers = []int{7, 9, 5, 12, 1, 3}
	fmt.Println("Before insertion sort:", numbers)
	sort(numbers)
	fmt.Println("After insertion sort:", numbers)
}
