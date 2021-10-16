package main

import "fmt"

type Data struct {
	Input  []int
	Output [][]int
}

func (data *Data) Permutation() {
	data.permut(len(data.Input), data.Input)
}

func (data *Data) permut(length int, aux []int) {
	var copyOfCurrentAux []int

	if length == 1 {
		copyOfCurrentAux = append([]int(nil), aux...)
		data.Output = append(data.Output, copyOfCurrentAux)
	}

	j := 0
	for i := 0; i < length; i++ {
		data.permut(length-1, aux)
		if length%2 == 0 {
			j = i
		}

		aux[j], aux[length-1] = aux[length-1], aux[j]
	}
}

func main() {
	data := Data{Input: []int{1, 4, 2}}
	data.Permutation()

	fmt.Println(data.Output)
}
